package request

type TransactionCreateReq struct {
	Items []TransactionItemReq `json:"items"`
}

type TransactionItemReq struct {
	ProductID int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type TransactionUpdateStatusReq struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}
