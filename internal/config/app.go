package config

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)


type BootstrapConfig struct {
	DB *gorm.DB
	App *gin.Engine
	Log *logrus.Logger
	Validator *validator.Validate
	Config *viper.Viper
}

func NewServer(config *BootstrapConfig) *BootstrapConfig{
	return &BootstrapConfig{
		DB: config.DB,
		App: config.App,
		Log: config.Log,
		Validator: config.Validator,
		Config: config.Config,
	}
}