package service

import (
	"context"

	"test/integration"
	integrationDto "test/integration/transactionDto"
	"test/service/dto"
)

type BoardCastService interface {
	BoardCastTransaction(ctx context.Context, request dto.BoardCastRequestDto) (*dto.BoardCastResponseDto, error)
	UtilizeTransaction(ctx context.Context, request string) (*dto.UtilizeResponseDto, error)
}

type boardCastService struct {
	transactionIntegration integration.TransactionIntegration
}

func NewBoardCastService(transactionIntegration integration.TransactionIntegration) BoardCastService {
	return &boardCastService{
		transactionIntegration: transactionIntegration,
	}
}

func (service *boardCastService) BoardCastTransaction(ctx context.Context, request dto.BoardCastRequestDto) (*dto.BoardCastResponseDto, error) {
	transaction := integrationDto.TransactionRequestDto{
		Symbol:    request.Symbol,
		Price:     request.Price,
		TimeStamp: request.TimeStamp,
	}

	data, err := service.transactionIntegration.BoardCastTransaction(ctx, transaction)
	if err != nil {
		panic(err)
	}

	response := dto.BoardCastResponseDto{
		TxHash: data.Tx_hash,
	}

	return &response, nil

}

func (service *boardCastService) UtilizeTransaction(ctx context.Context, request string) (*dto.UtilizeResponseDto, error) {

	data, err := service.transactionIntegration.UtilizeTransaction(ctx, request)
	if err != nil {
		panic(err)
	}

	response := dto.UtilizeResponseDto{}

	switch data.TxStatus {
	case "CONFIRMED": // Should be enum
		response.TxStatus = "CONFIRMED"
		response.Description = "Transaction has been processed and confirmed"
	case "FAILED": // Should be enum
		response.TxStatus = "FAILED"
		response.Description = "Transaction failed to process"
	case "PENDING": // Should be enum
		response.TxStatus = "PENDING"
		response.Description = "Transaction is awaiting processing"
	case "DNE": // Should be enum
		response.TxStatus = "DNE"
		response.Description = "Transaction does not exist"
	}

	return &response, nil

}
