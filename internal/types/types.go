package types

import (
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}
