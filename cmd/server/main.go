package main

import (
	"fmt"
	"net/http"
	"github.com/dettarune/goTokoo/internal/config"
)


func main(){

	host := config.NewViper().GetString("server.host")
	port := config.NewViper().GetString("server.port")
	address := fmt.Sprintf("%s:%s", host, port)


	viperConfig := config.NewViper()
	logrus := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig, logrus)
	validator := config.NewValidator(viperConfig)


	bootstrap := &config.BootstrapConfig{
		DB: db,
		Log: logrus,
		Validator: validator,
		Config: viperConfig,
	}

	app := config.NewServer(bootstrap)

	logrus.Infof("Server Running at http://%s", address)
	err := http.ListenAndServe(address, app)
	if err != nil {
		logrus.Fatalf("Failed To Start Server: %v", err)
	}
	
}