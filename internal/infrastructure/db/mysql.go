package db

import (
	"fmt"

	"github.com/tranChiVi0123/modular_monolithic_clean_architecture/internal/infrastructure/config"
	gormMySQLDriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func BuildMysqlConnection(config config.EnvConfigs) (*gorm.DB, error) {
	// BuildConnection is a function that will return a connection to the database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.FlezzDBUsername,
		config.FlezzDBPassword,
		config.FlezzDBHost,
		config.FlezzDBPort,
		config.FlezzDBDatabase,
	)

	db, err := gorm.Open(gormMySQLDriver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
