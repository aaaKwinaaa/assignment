package dto

type BoardCastRequestDto struct {
	Symbol    string `json:"symbol" binding:"required"`
	Price     int    `json:"price"  binding:"required"`
	TimeStamp int    `json:"timeStamp"  binding:"required"`
}
