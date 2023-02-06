package party

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"fmt"
	"math/big"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func generateEVMAccount() *ecdsa.PrivateKey {
	privateKey, _ := crypto.GenerateKey()
	return privateKey
}

func (am AppModule) waitAndVerifyEVMChain(ctx sdk.Context, client, client2 *ethclient.Client, request AccountWatchRequest) {
	// awrr := &AccountWatchRequestResult{
	// 	AccountWatchRequest: request,
	// 	Result:              OUTCOME_SUCCESS,
	// }
	// am.wg.Add(1)
	// // sleep for a random ammount of time between 5 and 10 seconds
	// time.Sleep(time.Duration(rand.Intn(5)+5) * time.Second)
	// am.dispatch(ctx, awrr)
	// am.wg.Wait()
	// return

	// create a ticker that ticks every 30 seconds
	// ticker := time.NewTicker(time.Second * 30)

	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

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
				fmt.Println("error getting balance: " + err.Error())
				continue
			}
			// if the balance is equal to the amount, verify with the
			// second RPC server.
			fmt.Printf("balance of %s is %s. looking for ammount %s", request.Account, balance.String(), request.Amount.String())
			if balance.Cmp(request.Amount) == 0 || balance.Cmp(request.Amount) == 1 {
				verifiedBalance, err := client2.BalanceAt(context.Background(), account, nil)
				if err != nil {
					fmt.Println("error getting balance: " + err.Error())
					continue
				}

				if verifiedBalance.Cmp(request.Amount) == 0 || verifiedBalance.Cmp(request.Amount) == 1 {
					// send a complete order event
					awrr := &AccountWatchRequestResult{
						AccountWatchRequest: request,
						Result:              OUTCOME_SUCCESS,
					}
					fmt.Printf("dispatching %s", awrr.Result)
					am.wg.Add(1)
					am.dispatch(ctx, awrr)
					am.wg.Wait()
					canILive = false
					return
				} else {
					continue
				}
			}
		case <-timer.C:
			fmt.Printf("account: %s timed out", request.Account)
			awrr := &AccountWatchRequestResult{
				AccountWatchRequest: request,
				Result:              OUTCOME_TIMEOUT,
			}
			am.wg.Add(1)
			am.dispatch(ctx, awrr)
			am.wg.Wait()
			canILive = false
			return
		}
	}
}

func (am AppModule) sendCoreEVMAsset(walletPrivK, walletPubK, toAddress string, amount *big.Int, txid string, rpcClient *ethclient.Client) error {
	// view the current balance of the paying wallet
	ecpk := ecdsa.PublicKey{}
	ecpk.X, ecpk.Y = elliptic.Unmarshal(crypto.S256(), common.FromHex(walletPubK))
	account := crypto.PubkeyToAddress(ecpk)
	// send the currency to the buyer
	// read nonce
	nonce, err := rpcClient.PendingNonceAt(context.Background(), account)
	if err != nil {
		return err
	}

	// create gas params
	gasLimit := uint64(31000) // in units
	gasPrice, err := rpcClient.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	// convert the string address to an address
	qualifiedAddress := common.HexToAddress(toAddress)

	// create a transaction
	tx := ethTypes.NewTransaction(nonce, qualifiedAddress, amount, gasLimit, gasPrice, nil)

	// fetch chain id
	chainID, err := rpcClient.NetworkID(context.Background())
	if err != nil {
		return err
	}

	//convert from a private key in a string to a *ecdsa.PrivateKey
	fromWallet, err := crypto.HexToECDSA(walletPrivK)
	if err != nil {
		return err
	}

	// sign the transaction
	signedTx, err := ethTypes.SignTx(tx, ethTypes.NewEIP155Signer(chainID), fromWallet)
	if err != nil {
		return err
	}

	// send the transaction
	err = rpcClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	return nil
}
