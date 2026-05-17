package repositories

import (
	"fmt"
	"rest-api-go/config"
	"rest-api-go/models"

	"gorm.io/gorm"
)

type AddressRepository struct {
	db *gorm.DB
}

func AddressRepositoryImpl(db *gorm.DB) *AddressRepository {
	if config.DB == nil {
		fmt.Println("DB is nil!")
	}
	return &AddressRepository{
		db: config.DB,
	}
}

func (r *AddressRepository) GetList() ([]models.Address, error) {
	var addresss []models.Address
	err := r.db.Order("name ASC").Find(&addresss).Error
	return addresss, err
}

func (r *AddressRepository) GetByID(addressId int64) (*models.Address, error) {
	var address models.Address
	err := r.db.Where("id = ?", addressId).First(&address).Error
	return &address, err
}

func (r *AddressRepository) Create(addresss *models.Address) (*models.Address, error) {
	if err := r.db.Save(addresss).Error; err != nil {
		return nil, err
	}
	return addresss, nil
}

func (r *AddressRepository) Update(addresss *models.Address) (*models.Address, error) {
	if err := r.db.Updates(addresss).Error; err != nil {
		return nil, err
	}
	return addresss, nil
}

func (r *AddressRepository) Delete(addressId int64) error {
	return r.db.Delete(&models.Address{}, "id = ?", addressId).Error
}
