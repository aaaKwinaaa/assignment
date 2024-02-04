package transactionDto

type TransactionRequestDto struct {
	Symbol    string `json:"symbol"`
	Price     int    `json:"price"`
	TimeStamp int    `json:"timestamp"`
}
