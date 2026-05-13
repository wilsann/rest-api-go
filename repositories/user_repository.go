package repositories

import (
	"fmt"
	"rest-api-go/config"
	"rest-api-go/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func UserRepositoryImpl(db *gorm.DB) *UserRepository {
	if config.DB == nil {
		fmt.Println("DB is nil!")
	}
	return &UserRepository{
		db: config.DB,
	}
}

func (r *UserRepository) GetList() (users []models.User, err error) {
	err = r.db.Order("name ASC").Find(&users).Error
	return users, err
}

func (r *UserRepository) GetByID(userId int64) (*models.User, error) {
	var users models.User
	err := r.db.Where("id = ?", userId).Error
	return &users, err
}

func (r *UserRepository) Create(users *models.User) (*models.User, error) {
	if err := r.db.Save(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Update(users *models.User) (*models.User, error) {
	if err := r.db.Updates(users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) Delete(userId int64) error {
	return r.db.Delete(&models.User{}, "id = ?", userId).Error
}
