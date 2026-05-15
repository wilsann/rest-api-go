package usecase

import (
	"errors"
	"rest-api-go/domain/request"
	"rest-api-go/domain/response"
	"rest-api-go/models"
	"rest-api-go/repositories"
	"time"
)

type TransactionUsecase interface {
	TransactionList(userId int64) ([]response.TransactionListResponse, error)
	TransactionDetail(transactionId int64) ([]response.TransactionDetailDataResponse, error)
	TransactionCreate(req request.TransactionCreateReq, userId int64) ([]response.TransactionDetailResponse, error)
	TransactionUpdateStatus(req request.TransactionUpdateStatusReq, userId int64) (*response.TransactionListResponse, error)
	TransactionDelete(productId int64) error
}

type TransactionUsecaseInteractor struct {
	TransactionRepository       repositories.TransactionRepository
	TransactionDetailRepository repositories.TransactionDetailRepository
	ProductRepository           repositories.ProductRepository
}

func TransactionUsecaseImpl(
	transactionRepository repositories.TransactionRepository,
	transactionDetailRepository repositories.TransactionDetailRepository,
	productRepository repositories.ProductRepository,
) TransactionUsecase {
	return &TransactionUsecaseInteractor{
		TransactionRepository:       transactionRepository,
		TransactionDetailRepository: transactionDetailRepository,
		ProductRepository:           productRepository,
	}
}

func (u *TransactionUsecaseInteractor) TransactionList(userId int64) ([]response.TransactionListResponse, error) {
	transactions, err := u.TransactionRepository.GetList(userId)
	if err != nil {
		return nil, errors.New("Failed get transaction list.")
	}

	var transactionsList []response.TransactionListResponse
	for _, p := range transactions {
		transactionsList = append(transactionsList, response.TransactionListResponse{
			ID:          p.ID,
			UserID:      p.UserID,
			TotalAmount: p.TotalAmount,
			Status:      p.Status,
			CreatedAt:   p.CreatedAt,
		})
	}

	return transactionsList, nil
}

func (u *TransactionUsecaseInteractor) TransactionDetail(transactionId int64) ([]response.TransactionDetailDataResponse, error) {
	transactions, err := u.TransactionDetailRepository.GetByTransactionID(transactionId)
	if err != nil {
		return nil, errors.New("Failed get data by transaction ID")
	}

	var transactionDetailList []response.TransactionDetailDataResponse
	for _, td := range transactions {
		transactionDetailList = append(transactionDetailList, response.TransactionDetailDataResponse{
			ID:              td.ID,
			TransactionID:   td.TransactionID,
			ProductID:       td.ProductID,
			Quantity:        td.Quantity,
			PriceAtPurchase: td.PriceAtPurchase,
			SubTotal:        td.SubTotal,
		})
	}

	return transactionDetailList, nil
}

func (u *TransactionUsecaseInteractor) TransactionCreate(req request.TransactionCreateReq, userId int64) ([]response.TransactionDetailResponse, error) {
	transaction := models.Transaction{
		UserID:      userId,
		TotalAmount: 0,
		Status:      "pending",
		CreatedAt:   time.Now(),
	}

	createdTransaction, err := u.TransactionRepository.Create(&transaction)
	if err != nil {
		return nil, errors.New("Failed create new transaction.")
	}

	var totalAmount float64
	transactionDetails := []models.TransactionDetail{}
	detailResponses := []response.TransactionDetailResponse{}
	detailDataResp := []response.TransactionDetailDataResponse{}

	for _, item := range req.Items {
		product, err := u.ProductRepository.GetByID(item.ProductID)
		if err != nil {
			return nil, errors.New("Failed get product by ID.")
		}

		subTotal := float64(item.Quantity) * product.Price
		totalAmount += subTotal

		details := models.TransactionDetail{
			TransactionID:   createdTransaction.ID,
			ProductID:       item.ProductID,
			Quantity:        item.Quantity,
			PriceAtPurchase: product.Price,
			SubTotal:        subTotal,
			CreatedAt:       time.Now(),
		}
		transactionDetails = append(transactionDetails, details)
	}

	createdDetails, err := u.TransactionDetailRepository.CreateBulk(transactionDetails)
	if err != nil {
		return nil, errors.New("Failed bulk create transaction detail data.")
	}

	err = u.TransactionRepository.UpdateTotalAmount(totalAmount, createdTransaction.ID)
	if err != nil {
		return nil, errors.New("Failed update total amount for transaction.")
	}

	for _, detail := range createdDetails {
		detailDataResp = append(detailDataResp, response.TransactionDetailDataResponse{
			ID:              detail.ID,
			TransactionID:   detail.TransactionID,
			ProductID:       detail.ProductID,
			Quantity:        detail.Quantity,
			PriceAtPurchase: detail.PriceAtPurchase,
			SubTotal:        detail.SubTotal,
		})
	}

	detailResponses = append(detailResponses, response.TransactionDetailResponse{
		ID:                createdTransaction.ID,
		UserID:            createdTransaction.UserID,
		TotalAmount:       createdTransaction.TotalAmount,
		Status:            createdTransaction.Status,
		TransactionDetail: detailDataResp,
	})

	return detailResponses, nil
}

func (u *TransactionUsecaseInteractor) TransactionUpdateStatus(req request.TransactionUpdateStatusReq, userId int64) (*response.TransactionListResponse, error) {
	transaction, err := u.TransactionRepository.GetByID(req.ID)
	if err != nil {
		return nil, errors.New("Failed fetch transaction by ID.")
	}

	transaction.Status = req.Status

	data, err := u.TransactionRepository.Update(transaction)
	if err != nil {
		return nil, errors.New("Failed update transaction.")
	}

	return &response.TransactionListResponse{
		ID:          data.ID,
		UserID:      data.UserID,
		TotalAmount: data.TotalAmount,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
	}, nil
}

func (u *TransactionUsecaseInteractor) TransactionDelete(productId int64) error {
	transaction, err := u.TransactionRepository.GetByID(productId)
	if err != nil {
		return errors.New("Failed get transaction by ID")
	}

	details, err := u.TransactionDetailRepository.GetByTransactionID(transaction.ID)
	if err != nil {
		return errors.New("Failed get transaction details by transactions ID")
	}

	for _, detail := range details {
		err = u.TransactionDetailRepository.Delete(detail.ID)
		if err != nil {
			return errors.New("Failed delete transaction details by transactions ID")
		}
	}

	err = u.TransactionRepository.Delete(productId)
	if err != nil {
		return errors.New("Failed delete transaction by ID.")
	}

	return err
}
