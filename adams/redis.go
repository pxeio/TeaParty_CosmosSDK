package adams

func (e *ExchangeServer) StoreSellOrders(orders []SellOrder) error {
	// store the order in the database
	// json, err := json.Marshal(orders)
	// if err != nil {
	// 	return err
	// }

	// // store the order in the database
	// err = e.redisClient.Set(context.Background(), "orders", json, 0).Err()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (e *ExchangeServer) updateOrdersInDB(orders []SellOrder) error {
	// store the order in the database
	// json, err := json.Marshal(orders)
	// if err != nil {
	// 	return err
	// }

	// err = e.redisClient.Set(context.Background(), "orders", json, 0).Err()
	// if err != nil {
	// 	return err
	// }

	return nil
}

// func (e *ExchangeServer) completeOrdersInDB(orders []CompletedOrder) error {
// 	// store the order in the database
// 	// ojs, err := json.Marshal(orders)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	err := e.redisClient.Set(context.Background(), "completeorders", orders, 0).Err()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (e *ExchangeServer) updateCompleteOrdersInDB(orders []CompletedOrder) error {
	// store the order in the database
	// ojs, err := json.Marshal(orders)
	// if err != nil {
	// 	return err
	// }

	// err = e.redisClient.Set(context.Background(), "completeorders", ojs, 0).Err()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (e *ExchangeServer) updateFailedOrdersInDB(orders CompletedOrder) error {
	// store the order in the database
	// ojs, err := json.Marshal(orders)
	// if err != nil {
	// 	return err
	// }

	// err = e.redisClient.Set(context.Background(), "failedorders", ojs, 0).Err()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func (e *ExchangeServer) fetchOrdersFromDB() ([]SellOrder, error) {
	// fetch the orders from the database
	// orders, err := e.redisClient.Get(context.Background(), "orders").Result()
	// if err != nil {
	// 	return nil, err
	// }

	// // unmarshal the orders
	// var o []SellOrder
	// err = json.Unmarshal([]byte(orders), &o)
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
