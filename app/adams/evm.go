package adams

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (e *ExchangeServer) generateEVMAccount(chain string) *ecdsa.PrivateKey {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	pk := hexutil.Encode(privateKeyBytes)[2:]
	e.logger.Debug("Generated " + chain + " Private Key: " + pk)
	return privateKey
}

func (a *ExchangeServer) waitAndVerifyEVMChain(ctx context.Context, client, client2 *ethclient.Client, request AccountWatchRequest) {
	if !a.watch {
		a.logger.Info("dev mode is on, not watching for payment. Returning success")
		awrr := &AccountWatchRequestResult{
			AccountWatchRequest: request,
			Result:              "suceess",
		}

		if err := a.Dispatch(awrr); err != nil {
			a.logger.Error("error dispatching account watch request result: " + err.Error())
		}
		return
	}
	a.logger.Info("Watching for " + request.Account + " to have a payment of " + request.Amount.String() + " on chain " + request.Chain)

	// create a ticker that ticks every 30 seconds
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()
	if a.dev {
		ticker = time.NewTicker(time.Second * 10)
	}

	// create a timer that times out after the specified timeout
	timer := time.NewTimer(time.Second * time.Duration(request.TimeOut))
	defer timer.Stop()

	account := common.HexToAddress(request.Account)

	// start a for loop that checks the balance of the address
	canILive := true
	for canILive {
		select {
		case <-ticker.C:
			balance, err := client.BalanceAt(context.Background(), account, nil)
			if err != nil {
				a.logger.Error("occured getting balance of " + request.Account + ": " + err.Error())
				return
			}
			a.logger.Infof("balance of %v is %v on chain %v", account, balance, request.Chain)
			// if the balance is equal to the amount, verify with the
			// second RPC server.
			if balance.Cmp(request.Amount) == 0 || balance.Cmp(request.Amount) == 1 {
				verifiedBalance, err := client2.BalanceAt(context.Background(), account, nil)
				if err != nil {
					a.logger.Error("occured getting balance of " + request.Account + ": " + err.Error() + " from the secondary ETH RPC server")
					return
				}

				if verifiedBalance.Cmp(request.Amount) == 0 || verifiedBalance.Cmp(request.Amount) == 1 {
					a.logger.Info("attempting to complete order " + request.TransactionID)
					// send a complete order event
					awrr := &AccountWatchRequestResult{
						AccountWatchRequest: request,
						Result:              "suceess",
					}

					if err := a.Dispatch(awrr); err != nil {
						a.logger.Error("error dispatching account watch request result: " + err.Error())
					}
					canILive = false
					return
				} else {
					a.logger.Error("balance of " + request.Account + " is not equal to " + request.Amount.String())
					return
				}
			}
		case <-timer.C:
			// if the timer times out, return an error
			e := fmt.Sprintf("timeout occured waiting for " + request.Account + " to have a payment of " + request.Amount.String())
			a.logger.Info(e)
			awrr := &AccountWatchRequestResult{
				AccountWatchRequest: request,
				Result:              "error",
			}

			if err := a.Dispatch(awrr); err != nil {
				a.logger.Error("error dispatching account watch request result: " + err.Error())
			}

			canILive = false

			return
		}
	}

}

func (e *ExchangeServer) sendCoreEVMAsset(fromWallet *ecdsa.PrivateKey, toAddress string, amount *big.Int, txid string, rpcClient *ethclient.Client) error {
	// view the current balance of the paying wallet
	account := crypto.PubkeyToAddress(fromWallet.PublicKey)
	// send the currency to the buyer
	// read nonce
	nonce, err := rpcClient.PendingNonceAt(context.Background(), account)
	if err != nil {
		e.logger.Error("cannot get nonce for " + account.String() + ": " + err.Error())
		return err
	}

	// create gas params
	gasLimit := uint64(31000) // in units
	gasPrice, err := rpcClient.SuggestGasPrice(context.Background())
	if err != nil {
		e.logger.Error("error getting gas price: " + err.Error())
		return err
	}

	// convert the string address to an address
	qualifiedAddress := common.HexToAddress(toAddress)

	// create a transaction
	tx := types.NewTransaction(nonce, qualifiedAddress, amount, gasLimit, gasPrice, nil)

	// fetch chain id
	chainID, err := rpcClient.NetworkID(context.Background())
	if err != nil {
		e.logger.Error("occured getting chain id: " + err.Error())
		return err
	}

	// sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), fromWallet)
	if err != nil {
		e.logger.Error("error signing transaction: " + err.Error())
		return err
	}

	// send the transaction
	err = rpcClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		e.logger.Error("error sending transaction: " + err.Error())
		return err
	}

	e.logger.Info("tx sent: " + signedTx.Hash().Hex() + "txid: " + txid)
	return nil
}
