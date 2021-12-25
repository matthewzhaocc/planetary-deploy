package types

import (
	"github.com/matthewzhaocc/planetary-deploy/internal/log"

	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Logger *log.Logger
}

type ServerConfig struct {
	SqlType   string `json:"engine"`
	SqlString string `json:"endpoint"`
}
