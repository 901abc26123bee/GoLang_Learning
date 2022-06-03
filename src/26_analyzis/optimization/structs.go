package optimize

type Request struct {
	TransactionID string `json:"translation_id"`
	PayLoad []int `json:"payload"`
}

type Response struct {
	TransactionID string `json:"transaction_id"`
	Expression string `json:"exp"`
}