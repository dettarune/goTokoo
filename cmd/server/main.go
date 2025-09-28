package main

import (
	"fmt"
	"net/http"
	"github.com/dettarune/goTokoo/internal/config"
	"github.com/dettarune/goTokoo/internal/infrastructure/router"
)


func main(){

	host := config.NewViper().GetString("server.host")
	port := config.NewViper().GetString("server.port")
	address := fmt.Sprintf("%s:%s", host, port)


	viperConfig := config.NewViper()
	logrus := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, logrus)
	validator := config.NewValidator(viperConfig)

	appConfig := &config.BootstrapConfig{
		DB: db,
		Log: logrus,
		Config: viperConfig,
		Validator: validator,
	}

	router := router.NewRouter(appConfig)

	logrus.Infof("Server Running at http://%s", address)
	err := http.ListenAndServe(address, router)
	if err != nil {
		logrus.Fatalf("Failed To Start Server: %v", err)
	}
	
}