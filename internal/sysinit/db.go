package sysinit

import (
	"errors"

	"github.com/matthewzhaocc/planetary-deploy/internal/types"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func GetDB(config types.ServerConfig) (*gorm.DB, error) {
	if config.SqlType == "mysql" {
		db, err := gorm.Open(mysql.Open(config.SqlString), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	if config.SqlType == "postgres" {
		db, err := gorm.Open(postgres.Open(config.SqlType), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	if config.SqlType == "sqlite" {
		db, err := gorm.Open(sqlite.Open(config.SqlType), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	if config.SqlType == "mssql" {
		db, err := gorm.Open(sqlserver.Open(config.SqlString), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}

	return nil, errors.New("Invalid DB Type.")
}
