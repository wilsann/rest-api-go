package repositories

import (
	"fmt"
	"rest-api-go/config"
	"rest-api-go/models"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func TransactionRepositoryImpl(db *gorm.DB) *TransactionRepository {
	if config.DB == nil {
		fmt.Println("DB is nil!")
	}
	return &TransactionRepository{
		db: config.DB,
	}
}

func (r *TransactionRepository) GetList(userId int64) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Where("user_id = ?", userId).Order("created_at ASC").Find(&transactions).Error
	return transactions, err
}

func (r *TransactionRepository) GetByID(transactionId int64) (*models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Where("id = ?", transactionId).First(&transaction).Error
	return &transaction, err
}

func (r *TransactionRepository) Create(transactions *models.Transaction) (*models.Transaction, error) {
	if err := r.db.Save(transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) Update(transactions *models.Transaction) (*models.Transaction, error) {
	if err := r.db.Updates(transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *TransactionRepository) UpdateTotalAmount(totalAmount float64, transactionID int64) error {
	if err := r.db.Debug().
		Model(&models.Transaction{}).
		Where("id = ?", transactionID).
		Update("total_amount", totalAmount).Error; err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) Delete(transactionId int64) error {
	return r.db.Delete(&models.Transaction{}, "id = ?", transactionId).Error
}
