package repositories

import (
	"fmt"
	"rest-api-go/config"
	"rest-api-go/models"

	"gorm.io/gorm"
)

type TransactionDetailRepository struct {
	db *gorm.DB
}

func TransactionDetailRepositoryImpl(db *gorm.DB) *TransactionDetailRepository {
	if config.DB == nil {
		fmt.Println("DB is nil!")
	}
	return &TransactionDetailRepository{
		db: config.DB,
	}
}

func (r *TransactionDetailRepository) GetList() ([]models.TransactionDetail, error) {
	var transactionDetails []models.TransactionDetail
	err := r.db.Order("name ASC").Find(&transactionDetails).Error
	return transactionDetails, err
}

func (r *TransactionDetailRepository) GetByID(transactionDetailId int64) (*models.TransactionDetail, error) {
	var transactionDetail models.TransactionDetail
	err := r.db.Where("id = ?", transactionDetailId).First(&transactionDetail).Error
	return &transactionDetail, err
}

func (r *TransactionDetailRepository) GetByTransactionID(transactionId int64) ([]models.TransactionDetail, error) {
	var transactionDetails []models.TransactionDetail
	err := r.db.
		Model(&models.TransactionDetail{}).
		Joins("INNER JOIN transactions t ON t.id = transaction_details.transaction_id").
		Where("t.id = ?", transactionId).
		Scan(&transactionDetails).Error

	return transactionDetails, err
}

func (r *TransactionDetailRepository) Create(transactionDetails *models.TransactionDetail) (*models.TransactionDetail, error) {
	if err := r.db.Save(transactionDetails).Error; err != nil {
		return nil, err
	}
	return transactionDetails, nil
}

func (r *TransactionDetailRepository) CreateBulk(transactionDetails []models.TransactionDetail) ([]models.TransactionDetail, error) {
	if err := r.db.Create(&transactionDetails).Error; err != nil {
		return nil, err
	}
	return transactionDetails, nil
}

func (r *TransactionDetailRepository) Update(transactionDetails *models.TransactionDetail) (*models.TransactionDetail, error) {
	if err := r.db.Updates(transactionDetails).Error; err != nil {
		return nil, err
	}
	return transactionDetails, nil
}

func (r *TransactionDetailRepository) Delete(transactionDetailId int64) error {
	return r.db.Delete(&models.TransactionDetail{}, "id = ?", transactionDetailId).Error
}
