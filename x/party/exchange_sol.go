package party

import (
	"context"
	"math/big"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	solRPC "github.com/gagliardetto/solana-go/rpc"
)

func createSolanaAccount() AccountResponse {
	account := solana.NewWallet()
	ar := &AccountResponse{
		PrivateKey: account.PrivateKey.String(),
		PublicKey:  account.PublicKey().String(),
	}
	return *ar
}

func (am AppModule) waitAndVerifySOLChain(ctx sdk.Context, request AccountWatchRequest, rpcClient, rpcClientTwo *solRPC.Client) error {
	// awrr := &AccountWatchRequestResult{
	// 	AccountWatchRequest: request,
	// 	Result:              OUTCOME_SUCCESS,
	// }
	// am.wg.Add(1)
	// am.dispatch(ctx, awrr)
	// am.wg.Wait()
	// return nil

	// the request.Amount is currently in ETH big.Int format convert to uint64
	amount, err := strconv.ParseUint(request.Amount.String(), 10, 64)
	if err != nil {
		return err
	}

	// convert from wei to lamports
	amount = amount / 1000000000

	// create a ticker that ticks every 30 seconds
	// ticker := time.NewTicker(time.Second * 30)
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	// create a timer that times out after the specified timeout
	timer := time.NewTimer(time.Second * time.Duration(request.TimeOut))
	defer timer.Stop()
	// start a for loop that checks the balance of the address
	canILive := true
	for canILive {
		select {
		case <-ticker.C:
			// create new solana public key from string
			pk, err := solana.PublicKeyFromBase58(request.Account)
			if err != nil {
				break
			}

			balance, err := rpcClient.GetBalance(context.Background(), pk, solRPC.CommitmentFinalized)
			if err != nil {
				break
			}

			// if the balance is equal to the amount, verify with the
			// second RPC server.
			if balance.Value >= amount {
				verifiedBalance, err := rpcClientTwo.GetBalance(context.Background(), pk, solRPC.CommitmentFinalized)
				if err != nil {
					break
				}

				if verifiedBalance.Value >= amount {
					// send a complete order event
					awrr := &AccountWatchRequestResult{
						AccountWatchRequest: request,
						Result:              OUTCOME_SUCCESS,
					}
					am.wg.Add(1)
					am.dispatch(ctx, awrr)
					am.wg.Wait()
					canILive = false
					return nil
				} else {
					break
				}
			}
		case <-timer.C:
			// if the timer times out, return an error
			awrr := &AccountWatchRequestResult{
				AccountWatchRequest: request,
				Result:              OUTCOME_TIMEOUT,
			}

			am.wg.Add(1)
			am.dispatch(ctx, awrr)
			am.wg.Wait()
			canILive = false
			return nil
		}
	}
	return nil
}

func (am AppModule) sendCoreSOLAsset(fromWalletPrivateKey, toAddress, txid string, amount *big.Int, rpcClient *solRPC.Client) error {
	privateKey, err := solana.PrivateKeyFromBase58(fromWalletPrivateKey)
	if err != nil {
		return err
	}

	toAddressPublicKey, err := solana.PublicKeyFromBase58(toAddress)
	if err != nil {
		return err
	}

	recent, err := rpcClient.GetRecentBlockhash(context.Background(), solRPC.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	// big.int to lamports
	amountLamparts := amount.Mul(amount, big.NewInt(1000000000))

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			system.NewTransferInstruction(
				amountLamparts.Uint64(),
				privateKey.PublicKey(),
				toAddressPublicKey,
			).Build(),
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(privateKey.PublicKey()),
	)
	if err != nil {
		return err
	}

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if privateKey.PublicKey().Equals(key) {
				return &privateKey
			}
			return nil
		},
	)
	if err != nil {
		return err
	}

	// TODO: Migrate to ws client so we can use the sendandconfirmtransaction method
	// Send transaction, and wait for confirmation:
	opts := solRPC.TransactionOpts{}
	_, err = rpcClient.SendTransactionWithOpts(
		context.Background(),
		tx,
		opts,
	)
	if err != nil {
		return err
	}
	return nil
}
