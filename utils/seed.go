package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"rest-api-go/models"

	"gorm.io/gorm"
)

func SeedSingleFile[T any](db *gorm.DB, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var items []T

	err = json.Unmarshal(data, &items)
	if err != nil {
		return err
	}

	result := db.Create(&items)
	if result.Error != nil {
		return result.Error
	}

	fmt.Println("Seeded:", path)

	return nil
}

func SeedFolder[T any](db *gorm.DB, folderPath string) error {
	files, err := os.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fullPath := filepath.Join(folderPath, file.Name())

		data, err := os.ReadFile(fullPath)
		if err != nil {
			return err
		}

		var items []T

		err = json.Unmarshal(data, &items)
		if err != nil {
			return err
		}

		result := db.Create(&items)
		if result.Error != nil {
			return result.Error
		}

		fmt.Println("Seeded:", fullPath)
	}

	return nil
}

func SeedAll(db *gorm.DB) error {
	if err := SeedSingleFile[models.Province](
		db,
		"./data/provinces/provinces.json",
	); err != nil {
		return err
	}

	if err := SeedFolder[models.Regency](
		db,
		"./data/regencies",
	); err != nil {
		return err
	}

	if err := SeedFolder[models.District](
		db,
		"./data/districts",
	); err != nil {
		return err
	}

	if err := SeedFolder[models.Village](
		db,
		"./data/villages",
	); err != nil {
		return err
	}

	return nil
}
