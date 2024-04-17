package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Open() (*gorm.DB, error) {
	dsn := "host=localhost user=itmo_rate dbname=itmo_rate_db port=5432"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
