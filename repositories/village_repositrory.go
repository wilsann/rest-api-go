package repositories

import (
	"fmt"
	"rest-api-go/config"
	"rest-api-go/models"

	"gorm.io/gorm"
)

type VillageRepository struct {
	db *gorm.DB
}

func VillageRepositoryImpl(db *gorm.DB) *VillageRepository {
	if config.DB == nil {
		fmt.Println("DB is nil!")
	}
	return &VillageRepository{
		db: config.DB,
	}
}

func (r *VillageRepository) GetList() ([]models.Village, error) {
	var villages []models.Village
	err := r.db.Order("name ASC").Find(&villages).Error
	return villages, err
}

func (r *VillageRepository) GetByID(villageId int64) (*models.Village, error) {
	var village models.Village
	err := r.db.Where("id = ?", villageId).First(&village).Error
	return &village, err
}

func (r *VillageRepository) Create(villages *models.Village) (*models.Village, error) {
	if err := r.db.Save(villages).Error; err != nil {
		return nil, err
	}
	return villages, nil
}

func (r *VillageRepository) Update(villages *models.Village) (*models.Village, error) {
	if err := r.db.Updates(villages).Error; err != nil {
		return nil, err
	}
	return villages, nil
}

func (r *VillageRepository) Delete(villageId int64) error {
	return r.db.Delete(&models.Village{}, "id = ?", villageId).Error
}
