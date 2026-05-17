package repositories

import (
	"fmt"
	"rest-api-go/config"
	"rest-api-go/models"

	"gorm.io/gorm"
)

type RegencyRepository struct {
	db *gorm.DB
}

func RegencyRepositoryImpl(db *gorm.DB) *RegencyRepository {
	if config.DB == nil {
		fmt.Println("DB is nil!")
	}
	return &RegencyRepository{
		db: config.DB,
	}
}

func (r *RegencyRepository) GetList() ([]models.Regency, error) {
	var regencys []models.Regency
	err := r.db.Order("name ASC").Find(&regencys).Error
	return regencys, err
}

func (r *RegencyRepository) GetByID(regencyId int64) (*models.Regency, error) {
	var regency models.Regency
	err := r.db.Where("id = ?", regencyId).First(&regency).Error
	return &regency, err
}

func (r *RegencyRepository) Create(regencys *models.Regency) (*models.Regency, error) {
	if err := r.db.Save(regencys).Error; err != nil {
		return nil, err
	}
	return regencys, nil
}

func (r *RegencyRepository) Update(regencys *models.Regency) (*models.Regency, error) {
	if err := r.db.Updates(regencys).Error; err != nil {
		return nil, err
	}
	return regencys, nil
}

func (r *RegencyRepository) Delete(regencyId int64) error {
	return r.db.Delete(&models.Regency{}, "id = ?", regencyId).Error
}
