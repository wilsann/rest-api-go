package repositories

import (
	"fmt"
	"rest-api-go/config"
	"rest-api-go/models"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func ProductRepositoryImpl(db *gorm.DB) *ProductRepository {
	if config.DB == nil {
		fmt.Println("DB is nil!")
	}
	return &ProductRepository{
		db: config.DB,
	}
}

func (r *ProductRepository) GetList() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Where("deleted_at IS NULL").Order("name ASC").Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetByID(productId int64) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ? AND deleted_at IS NULL", productId).First(&product).Error
	return &product, err
}

func (r *ProductRepository) Create(products *models.Product) (*models.Product, error) {
	if err := r.db.Save(products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Update(products *models.Product) (*models.Product, error) {
	if err := r.db.Updates(products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) Delete(productId int64) error {
	return r.db.Delete(&models.Product{}, "id = ?", productId).Error
}
