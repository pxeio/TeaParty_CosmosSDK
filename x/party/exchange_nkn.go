package party

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	partyTypes "github.com/TeaPartyCrypto/partychain/x/party/types"
	nkn "github.com/nknorg/nkn-sdk-go"
)

func notifySellerOfBuyer(co CompletedOrder) error {
	account, err := nkn.NewAccount(nil)
	if err != nil {
		return err
	}

	nknClient, err := nkn.NewMultiClient(account, "", 4, false, nil)
	if err != nil {
		return err
	}
	defer nknClient.Close()

	sn := &NKNNotification{
		Address: co.SellerEscrowWallet.PublicAddress,
		Amount:  co.Amount.String(),
		Network: co.Currency,
	}
	bytes, err := json.Marshal(sn)
	if err != nil {
		return err
	}

	<-nknClient.OnConnect.C
	onReply, err := nknClient.Send(nkn.NewStringArray(co.SellerNKNAddress), bytes, nil)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	select {
	case reply := <-onReply.C:
		switch string(reply.Data) {
		case "ok":
			return nil
		default:
			return fmt.Errorf("seller has encountered an error for order: " + co.OrderID)
		}
	case <-ctx.Done():
		return fmt.Errorf("seller has not responded to notification of new buyer for order: " + co.OrderID)
	}
}

func sendBuyerPayInfo(co CompletedOrder) error {
	account, err := nkn.NewAccount(nil)
	if err != nil {
		return err
	}
	nknClient, err := nkn.NewMultiClient(account, "", 4, false, nil)
	if err != nil {
		return err
	}
	defer nknClient.Close()

	sn := &NKNNotification{
		Address: co.BuyerEscrowWallet.PublicAddress,
		Amount:  co.Price.String(),
		Network: co.TradeAsset,
	}
	bytes, err := json.Marshal(sn)
	if err != nil {
		return err
	}

	<-nknClient.OnConnect.C
	onReply, err := nknClient.Send(nkn.NewStringArray(co.BuyerNKNAddress), bytes, nil)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	select {
	case reply := <-onReply.C:
		switch string(reply.Data) {
		case "ok":
			return nil
		default:
			return fmt.Errorf("buyer has encountered an error for order: " + co.OrderID)
		}
	case <-ctx.Done():
		return fmt.Errorf("buyer has not responded to notification of escrow information for order: " + co.OrderID)
	}
}

// sendPrivateKey is called to send the private key of an escrow wallet.
func (am AppModule) sendPrivateKey(order partyTypes.OrdersAwaitingFinalizer) error {
	fmt.Printf("sending private key to %s", order.NknAddress)
	account, err := nkn.NewAccount(nil)
	if err != nil {
		return err
	}

	// initialize the nkn client.
	nknClient, err := nkn.NewMultiClient(account, "", 4, false, nil)
	if err != nil {
		return err
	}

	ac := &NKNNotification{
		PrivateKey: order.WalletPrivateKey,
		Address:    order.WalletPublicKey,
		Chain:      order.Chain,
	}

	bytes, err := json.Marshal(ac)
	if err != nil {
		return err
	}

	<-nknClient.OnConnect.C
	onReply, err := nknClient.Send(nkn.NewStringArray(order.NknAddress), bytes, nil)
	if err != nil {
		return err
	}
	// create a timeout of 2 minutes
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	select {
	case reply := <-onReply.C:
		switch string(reply.Data) {
		case "ok":
			return nil
		default:
			return fmt.Errorf("buyer has encountered an error receiving sellers escrow wallet private key for order: " + order.Index)
		}
	case <-ctx.Done():
		return fmt.Errorf("buyer has not responded to notification of new PK for order: " + order.Index)
	}
	return nil
}
