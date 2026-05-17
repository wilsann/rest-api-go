package repositories

import (
	"fmt"
	"rest-api-go/config"
	"rest-api-go/models"

	"gorm.io/gorm"
)

type ProvinceRepository struct {
	db *gorm.DB
}

func ProvinceRepositoryImpl(db *gorm.DB) *ProvinceRepository {
	if config.DB == nil {
		fmt.Println("DB is nil!")
	}
	return &ProvinceRepository{
		db: config.DB,
	}
}

func (r *ProvinceRepository) GetList() ([]models.Province, error) {
	var provinces []models.Province
	err := r.db.Order("name ASC").Find(&provinces).Error
	return provinces, err
}

func (r *ProvinceRepository) GetByID(provinceId int64) (*models.Province, error) {
	var province models.Province
	err := r.db.Where("id = ?", provinceId).First(&province).Error
	return &province, err
}

func (r *ProvinceRepository) Create(provinces *models.Province) (*models.Province, error) {
	if err := r.db.Save(provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r *ProvinceRepository) Update(provinces *models.Province) (*models.Province, error) {
	if err := r.db.Updates(provinces).Error; err != nil {
		return nil, err
	}
	return provinces, nil
}

func (r *ProvinceRepository) Delete(provinceId int64) error {
	return r.db.Delete(&models.Province{}, "id = ?", provinceId).Error
}
