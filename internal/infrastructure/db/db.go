package db

import "gorm.io/gorm"

type DbConfig struct {
  // Main database connection in the future can split into multiple databases
  FDB *gorm.DB
}
