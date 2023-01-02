package adams

import (
	"math/big"
)

// func (e *ExchangeServer) Dispatch(awrr *AccountWatchRequestResult) error {
// 	if awrr.Result == "error" {
// 		e.logger.Infof("Account watch result: %s", awrr.Result)
// 		e.logger.Infof("closing the transaction: %s", awrr.AccountWatchRequest.TransactionID)
// 		for _, order := range e.completOrders {
// 			if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 				if order.BuyerPaymentComplete {
// 					e.logger.Infof("buyer funded the escrow wallet, refunding the buyer %s", order.BuyerNKNAddress)
// 					if err := e.refundBuyerViaEscrowWallet(order); err != nil {
// 						e.logger.Errorw("failed to refund buyer via escrow wallet...", err)
// 						if refundErr := e.sendREFUNDToBuyer(order); err != nil {
// 							e.logger.Errorw("failed to send refund to buyer on chain", err)
// 							err := e.updateFailedOrdersInDB(order)
// 							if err != nil {
// 								e.logger.Errorw("failed to update the failed orders in the db", err)
// 							}
// 							return refundErr
// 						}
// 					}

// 				}

// 				if order.SellerPaymentComplete {
// 					e.logger.Infof("seller funded the escrow wallet, refunding the seller %s", order.BuyerNKNAddress)
// 					if err := e.refundSellerViaEscrowWallet(order); err != nil {
// 						e.logger.Errorw("failed to refund seller", err)
// 						if refundErr := e.sendREFUNDToSeller(order); err != nil {
// 							e.logger.Errorw("failed to send refund to seller on chain", err)
// 							if err := e.updateFailedOrdersInDB(order); err != nil {
// 								e.logger.Errorw("failed to update the failed orders in the db", err)
// 							}
// 							return refundErr
// 						}
// 					}
// 				}
// 				e.closeOrder(&order)
// 				return nil
// 			}
// 		}

// 	}

// 	if awrr.Result == "suceess" {
// 		e.logger.Infof("Successfull Account watch result: %s", awrr.Result)
// 		if awrr.AccountWatchRequest.Seller {
// 			for i, order := range e.completOrders {
// 				if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 					e.completOrders[i].SellerPaymentComplete = true
// 					e.logger.Infof("seller payment complete for order %s", order.OrderID)
// 					if err := e.updateCompleteOrdersInDB(e.completOrders); err != nil {
// 						e.logger.Errorw("failed to update the complete orders in the db", err)
// 						return err
// 					}
// 					break
// 				}
// 			}
// 		} else {
// 			for i, order := range e.completOrders {
// 				if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 					e.logger.Infof("buyer payment complete for order %s", order.OrderID)
// 					e.completOrders[i].BuyerPaymentComplete = true
// 					if err := e.updateCompleteOrdersInDB(e.completOrders); err != nil {
// 						e.logger.Errorw("failed to update the complete orders in the db", err)
// 						return err
// 					}
// 					break
// 				}
// 			}
// 		}

// 		// check if the order is complete
// 		for _, order := range e.completOrders {
// 			if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 				if order.BuyerPaymentComplete && order.SellerPaymentComplete {
// 					e.logger.Infof("order %s is complete", order.OrderID)
// 					e.logger.Info("Attempting to find a matching order and complete it")
// 					var found bool
// 					for _, order := range e.completOrders {
// 						if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 							e.logger.Info("Found a matching order.. calling completeOrder")
// 							e.completeOrder(order)
// 							found = true
// 							return nil
// 						}
// 					}

// 					if !found {
// 						e.logger.Infof("No matching order found for transaction id: %s", awrr.AccountWatchRequest.TransactionID)
// 						return nil
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return nil
// }

func (e *ExchangeServer) Dispatch(awrr *AccountWatchRequestResult) error {
	// notify the party chain of the transaction outcome
	if err := notifyPartyChainOfWatchResult(e.partyNode, awrr); err != nil {
		e.logger.Errorw("failed to notify the party chain of the watch result", err)
		return err
	}

	return nil
}

// TODO:// move this logic to the Party chain
// 	if awrr.Result == "error" {
// 		e.logger.Infof("Account watch result: %s", awrr.Result)
// 		e.logger.Infof("closing the transaction: %s", awrr.AccountWatchRequest.TransactionID)

// 		for _, order := range e.completOrders {
// 			if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 				if order.BuyerPaymentComplete {
// 					e.logger.Infof("buyer funded the escrow wallet, refunding the buyer %s", order.BuyerNKNAddress)
// 					if err := e.refundBuyerViaEscrowWallet(order); err != nil {
// 						e.logger.Errorw("failed to refund buyer via escrow wallet...", err)
// 						if refundErr := e.sendREFUNDToBuyer(order); err != nil {
// 							e.logger.Errorw("failed to send refund to buyer on chain", err)
// 							err := e.updateFailedOrdersInDB(order)
// 							if err != nil {
// 								e.logger.Errorw("failed to update the failed orders in the db", err)
// 							}
// 							return refundErr
// 						}
// 					}

// 				}

// 				if order.SellerPaymentComplete {
// 					e.logger.Infof("seller funded the escrow wallet, refunding the seller %s", order.BuyerNKNAddress)
// 					if err := e.refundSellerViaEscrowWallet(order); err != nil {
// 						e.logger.Errorw("failed to refund seller", err)
// 						if refundErr := e.sendREFUNDToSeller(order); err != nil {
// 							e.logger.Errorw("failed to send refund to seller on chain", err)
// 							if err := e.updateFailedOrdersInDB(order); err != nil {
// 								e.logger.Errorw("failed to update the failed orders in the db", err)
// 							}
// 							return refundErr
// 						}
// 					}
// 				}
// 				e.closeOrder(&order)
// 				return nil
// 			}
// 		}

// 	}

// 	if awrr.Result == "suceess" {
// 		e.logger.Infof("Successfull Account watch result: %s", awrr.Result)
// 		if awrr.AccountWatchRequest.Seller {
// 			for i, order := range e.completOrders {
// 				if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 					e.completOrders[i].SellerPaymentComplete = true
// 					e.logger.Infof("seller payment complete for order %s", order.OrderID)
// 					if err := e.updateCompleteOrdersInDB(e.completOrders); err != nil {
// 						e.logger.Errorw("failed to update the complete orders in the db", err)
// 						return err
// 					}
// 					break
// 				}
// 			}
// 		} else {
// 			for i, order := range e.completOrders {
// 				if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 					e.logger.Infof("buyer payment complete for order %s", order.OrderID)
// 					e.completOrders[i].BuyerPaymentComplete = true
// 					if err := e.updateCompleteOrdersInDB(e.completOrders); err != nil {
// 						e.logger.Errorw("failed to update the complete orders in the db", err)
// 						return err
// 					}
// 					break
// 				}
// 			}
// 		}

// 		// check if the order is complete
// 		for _, order := range e.completOrders {
// 			if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 				if order.BuyerPaymentComplete && order.SellerPaymentComplete {
// 					e.logger.Infof("order %s is complete", order.OrderID)
// 					e.logger.Info("Attempting to find a matching order and complete it")
// 					var found bool
// 					for _, order := range e.completOrders {
// 						if order.OrderID == awrr.AccountWatchRequest.TransactionID {
// 							e.logger.Info("Found a matching order.. calling completeOrder")
// 							e.completeOrder(order)
// 							found = true
// 							return nil
// 						}
// 					}

// 					if !found {
// 						e.logger.Infof("No matching order found for transaction id: %s", awrr.AccountWatchRequest.TransactionID)
// 						return nil
// 					}
// 				}
// 			}
// 		}
// 	}

// 	return nil
// }

// closeOrder closes the order with the specified OrderID
// this is called after a sucessfull transaction
func (e *ExchangeServer) closeOrder(co *CompletedOrder) {
	for _, order := range e.orders {
		if order.TXID == co.OrderID {
			e.logger.Info("order: " + co.OrderID + " is closing")
			// e.orders = append(e.orders[:i], e.orders[i+1:]...)
			// err := e.updateOrdersInDB(e.orders)
			// if err != nil {
			// 	e.logger.Error("error updating the orders in the db: " + err.Error())
			// }

			// jsn, err := json.Marshal(co)
			// if err != nil {
			// 	e.logger.Error("error marshalling the completed orders: " + err.Error())
			// 	return
			// }

			// err = e.redisClient.Set(context.Background(), "completedorders", jsn, 0).Err()
			// if err != nil {
			// 	e.logger.Errorw("failed to save the complted orders map to redis")
			// }
			// return
		}
		e.logger.Info("order: " + co.OrderID + " not found")
	}
	return
}

// closeFailedOrder closes the order with the specified OrderID
// this is called after a failed transaction
func (e *ExchangeServer) closeFailedOrder(co *CompletedOrder) {
	for _, order := range e.orders {
		if order.TXID == co.OrderID {
			e.logger.Info("order: " + co.OrderID + " is closing")
			// e.orders = append(e.orders[:i], e.orders[i+1:]...)
			// err := e.updateOrdersInDB(e.orders)
			// if err != nil {
			// 	e.logger.Error("error updating the orders in the db: " + err.Error())
			// }

			// jsn, err := json.Marshal(co)
			// if err != nil {
			// 	e.logger.Error("error marshalling the completed orders: " + err.Error())
			// 	return
			// }

			// err = e.redisClient.Set(context.Background(), "completedorders", jsn, 0).Err()
			// if err != nil {
			// 	e.logger.Errorw("failed to save the complted orders map to redis")
			// }

			return
		}
		e.logger.Info("order: " + co.OrderID + " not found")
	}
	return
}

func (e *ExchangeServer) cancelOrder(OrderID string) {
	for _, order := range e.orders {
		if order.TXID == OrderID {
			// if err := e.refundSellerViaEscrowWallet(e.completOrders[i]); err != nil {
			// 	e.logger.Error("refund seller via escrow wallet failed, pushing the transaction to state")
			// }

			// if err := e.refundBuyerViaEscrowWallet(e.completOrders[i]); err != nil {
			// 	e.logger.Error("refund seller via escrow wallet failed, pushing the transaction to state")
			// }

			// e.logger.Info("order: " + OrderID + " is canceling")
			// e.orders = append(e.orders[:i], e.orders[i+1:]...)

			// // update the orders in the db
			// err := e.updateOrdersInDB(e.orders)
			// if err != nil {
			// 	e.logger.Error("error updating the orders in the db: " + err.Error())
			// }

			// // update the failed orders in the db
			// if err := e.updateFailedOrdersInDB(e.completOrders[i]); err != nil {
			// 	e.logger.Error("error updating the failed orders in the db: " + err.Error())
			// }

			return
		}
		e.logger.Info("order: " + OrderID + " not found")
	}
	return
}

func (e *ExchangeServer) completeOrder(order CompletedOrder) {
	e.logger.Info("order: " + order.OrderID + " is complete.. now completing the order")
	e.logger.Info("sending the buyer the private key for the escrow wallet")
	if err := e.sendSellerEscrowWalletPrivateKeyToBuyer(order); err != nil {
		e.logger.Error("error sending the buyer the private key for the escrow wallet: " + err.Error())
		e.logger.Info("attempting to send the buyer funds on chain")
		if err := e.sendFundsToBuyer(order); err != nil {
			e.logger.Error("error sending the funds to the buyer on chain: " + err.Error())
		}
	}

	e.logger.Info("sending the seller the private key for the escrow wallet")
	if err2 := e.sendBuyerEscrowWalletPrivateKeyToSeller(order); err2 != nil {
		e.logger.Error("error sending the buyer the private key for the escrow wallet: " + err2.Error())
		e.logger.Info("attempting to send the seller funds on chain")
		if err := e.sendFundsToSeller(order); err != nil {
			e.logger.Error("error sending the funds to the seller on chain: " + err.Error())
		}
	}

	e.closeOrder(&order)
}

func (e *ExchangeServer) sendFundsToBuyer(order CompletedOrder) error {
	// // currenty we do not support sending `kaspa` on chain
	// if order.Currency == KAS || order.Currency == BTC || order.Currency == RXD {
	// 	return fmt.Errorf("sending kaspa on chain is not supported")
	// }

	curencyFee := new(big.Int).Div(order.Amount, big.NewInt(100))
	order.Amount.Sub(order.Amount, curencyFee)
	assetFee := new(big.Int).Div(order.Price, big.NewInt(100))
	order.Price.Sub(order.Price, assetFee)

	// send the funds to the buyer
	switch order.Currency {
	case SOL:
		if err := e.SendCoreSOLAsset(order.SellerEscrowWallet.PrivateKey, order.BuyerShippingAddress, order.OrderID, order.Amount); err != nil {
			e.logger.Error("sending funds to buyer on SOL: " + err.Error())
			return err
		}
	case CEL:
		err := e.sendCoreEVMAsset(order.SellerEscrowWallet.ECDSA, order.BuyerShippingAddress, order.Amount, order.OrderID, e.celoNode.rpcClient)
		if err != nil {
			e.logger.Error("sending funds to buyer on CEL: " + err.Error())
			return err
		}

	case MO:
		err := e.sendCoreEVMAsset(order.SellerEscrowWallet.ECDSA, order.BuyerShippingAddress, order.Amount, order.OrderID, e.mineOnlium.rpcClient)
		if err != nil {
			e.logger.Error("sending funds to buyer on MO: " + err.Error())
			return err
		}

	case POL:
		err := e.sendCoreEVMAsset(order.SellerEscrowWallet.ECDSA, order.BuyerShippingAddress, order.Amount, order.OrderID, e.polygonNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to buyer on POL: " + err.Error())
			return err
		}

	case ETH:
		err := e.sendCoreEVMAsset(order.SellerEscrowWallet.ECDSA, order.BuyerShippingAddress, order.Amount, order.OrderID, e.ethNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to buyer on ETH: " + err.Error())
			return err
		}

	}
	return nil
}

func (e *ExchangeServer) sendREFUNDToBuyer(order CompletedOrder) error {
	// // currenty we do not support sending `kaspa` on chain
	// if order.Currency == KAS || order.Currency == BTC || order.Currency == RXD {
	// 	return fmt.Errorf("sending kaspa on chain is not supported")
	// }

	curencyFee := new(big.Int).Div(order.Amount, big.NewInt(100))
	order.Amount.Sub(order.Amount, curencyFee)
	assetFee := new(big.Int).Div(order.Price, big.NewInt(100))
	order.Price.Sub(order.Price, assetFee)

	// send the funds to the buyer
	switch order.Currency {
	case SOL:
		if err := e.SendCoreSOLAsset(order.BuyerEscrowWallet.PrivateKey, order.BuyerRefundAddress, order.OrderID, order.Amount); err != nil {
			e.logger.Error("sending funds to buyer on SOL: " + err.Error())
			return err
		}
	case CEL:
		err := e.sendCoreEVMAsset(order.BuyerEscrowWallet.ECDSA, order.BuyerRefundAddress, order.Price, order.OrderID, e.celoNode.rpcClient)
		if err != nil {
			e.logger.Error("sending funds to buyer on CEL: " + err.Error())
			return err
		}

	case MO:
		err := e.sendCoreEVMAsset(order.BuyerEscrowWallet.ECDSA, order.BuyerRefundAddress, order.Price, order.OrderID, e.mineOnlium.rpcClient)
		if err != nil {
			e.logger.Error("sending funds to buyer on MO: " + err.Error())
			return err
		}

	case POL:
		err := e.sendCoreEVMAsset(order.BuyerEscrowWallet.ECDSA, order.BuyerRefundAddress, order.Price, order.OrderID, e.polygonNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to buyer on POL: " + err.Error())
			return err
		}

	case ETH:
		err := e.sendCoreEVMAsset(order.BuyerEscrowWallet.ECDSA, order.BuyerRefundAddress, order.Price, order.OrderID, e.ethNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to buyer on ETH: " + err.Error())
			return err
		}

	}
	return nil
}

// sendFundsToSeller provides functionality to send funds to the seller
func (e *ExchangeServer) sendFundsToSeller(order CompletedOrder) error {
	// if order.TradeAsset == KAS || order.TradeAsset == BTC || order.TradeAsset == RXD {
	// 	return fmt.Errorf("sending kaspa on chain is not supported")
	// }

	curencyFee := new(big.Int).Div(order.Amount, big.NewInt(100))
	order.Amount.Sub(order.Amount, curencyFee)
	assetFee := new(big.Int).Div(order.Price, big.NewInt(100))
	order.Price.Sub(order.Price, assetFee)

	switch order.TradeAsset {
	case SOL:
		if err := e.SendCoreSOLAsset(order.BuyerEscrowWallet.PrivateKey, order.SellerShippingAddress, order.OrderID, order.Amount); err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}
	case CEL:
		err := e.sendCoreEVMAsset(order.BuyerEscrowWallet.ECDSA, order.SellerShippingAddress, order.Price, order.OrderID, e.celoNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}

	case MO:
		err := e.sendCoreEVMAsset(order.BuyerEscrowWallet.ECDSA, order.SellerShippingAddress, order.Price, order.OrderID, e.mineOnlium.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}

	case POL:
		err := e.sendCoreEVMAsset(order.BuyerEscrowWallet.ECDSA, order.SellerShippingAddress, order.Price, order.OrderID, e.polygonNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}

	case ETH:
		err := e.sendCoreEVMAsset(order.BuyerEscrowWallet.ECDSA, order.SellerShippingAddress, order.Price, order.OrderID, e.ethNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}

	}
	return nil
}

// sendREFUNDToSeller provides functionality to send funds to the seller
func (e *ExchangeServer) sendREFUNDToSeller(order CompletedOrder) error {
	// if order.TradeAsset == KAS || order.TradeAsset == BTC || order.TradeAsset == RXD {
	// 	return fmt.Errorf("sending kaspa on chain is not supported")
	// }

	curencyFee := new(big.Int).Div(order.Amount, big.NewInt(100))
	order.Amount.Sub(order.Amount, curencyFee)
	assetFee := new(big.Int).Div(order.Price, big.NewInt(100))
	order.Price.Sub(order.Price, assetFee)

	switch order.TradeAsset {
	case SOL:
		if err := e.SendCoreSOLAsset(order.SellerEscrowWallet.PrivateKey, order.SellerRefundAddress, order.OrderID, order.Amount); err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}

	case CEL:
		err := e.sendCoreEVMAsset(order.SellerEscrowWallet.ECDSA, order.SellerRefundAddress, order.Amount, order.OrderID, e.celoNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}

	case MO:
		err := e.sendCoreEVMAsset(order.SellerEscrowWallet.ECDSA, order.SellerRefundAddress, order.Amount, order.OrderID, e.mineOnlium.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}

	case POL:
		err := e.sendCoreEVMAsset(order.SellerEscrowWallet.ECDSA, order.SellerRefundAddress, order.Amount, order.OrderID, e.polygonNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}

	case ETH:
		err := e.sendCoreEVMAsset(order.SellerEscrowWallet.ECDSA, order.SellerRefundAddress, order.Amount, order.OrderID, e.ethNode.rpcClient)
		if err != nil {
			e.logger.Error("error sending funds to seller: " + err.Error())
			return err
		}

	}
	return nil
}

func FloatToBigInt(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	// Set precision if required.
	// bigval.SetPrec(64)

	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}
