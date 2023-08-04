package postgre

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
