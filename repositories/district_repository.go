package repositories

import (
	"fmt"
	"rest-api-go/config"
	"rest-api-go/models"

	"gorm.io/gorm"
)

type DistrictRepository struct {
	db *gorm.DB
}

func DistrictRepositoryImpl(db *gorm.DB) *DistrictRepository {
	if config.DB == nil {
		fmt.Println("DB is nil!")
	}
	return &DistrictRepository{
		db: config.DB,
	}
}

func (r *DistrictRepository) GetList() ([]models.District, error) {
	var districts []models.District
	err := r.db.Order("name ASC").Find(&districts).Error
	return districts, err
}

func (r *DistrictRepository) GetByID(districtId int64) (*models.District, error) {
	var district models.District
	err := r.db.Where("id = ?", districtId).First(&district).Error
	return &district, err
}

func (r *DistrictRepository) Create(districts *models.District) (*models.District, error) {
	if err := r.db.Save(districts).Error; err != nil {
		return nil, err
	}
	return districts, nil
}

func (r *DistrictRepository) Update(districts *models.District) (*models.District, error) {
	if err := r.db.Updates(districts).Error; err != nil {
		return nil, err
	}
	return districts, nil
}

func (r *DistrictRepository) Delete(districtId int64) error {
	return r.db.Delete(&models.District{}, "id = ?", districtId).Error
}
