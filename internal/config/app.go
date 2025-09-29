package config

import (
	httpDelivery "github.com/dettarune/goTokoo/internal/delivery/http"
	"github.com/dettarune/goTokoo/internal/delivery/router"
	"github.com/dettarune/goTokoo/internal/repository"
	"github.com/dettarune/goTokoo/internal/usecase"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB        *gorm.DB
	Log       *logrus.Logger
	Validator *validator.Validate
	Config    *viper.Viper
}

// NewServer build all dependencies and return chi router
func NewServer(config *BootstrapConfig) *chi.Mux {
	app := chi.NewRouter()

	// Repository
	userRepository := repository.NewUserRepository(config.Log)

	// UseCase
	userUseCase := usecase.NewUserUseCase(&usecase.UserUseCaseDeps{
		DB:            config.DB,
		Log:           config.Log,
		Validate:      config.Validator,
		UserRepository: userRepository,
	})

	// Controller
	userController := httpDelivery.NewUserController(userUseCase, config.Log)

	// Routing
	routeConfig := router.RouteConfig{
		App:            app,
		UserController: userController,
	}
	routeConfig.SetupRoute()

	return app
}
