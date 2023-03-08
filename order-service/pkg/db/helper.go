package db

func reduceCreateOrderTxParam(items []CreateOrderItemParam) []CreateOrderItemParam {
	orderMap := make(map[int64]CreateOrderItemParam)

	for _, item := range items {

		i, ok := orderMap[item.PID]
		if !ok {
			orderMap[item.PID] = item
		} else {
			i.Amount += item.Amount
			orderMap[item.PID] = i
		}
	}

	var newItems []CreateOrderItemParam
	for _, value := range orderMap {
		newItems = append(newItems, value)
	}
	return newItems
}
