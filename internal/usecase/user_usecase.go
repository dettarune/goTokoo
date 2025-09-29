package usecase

import (
	"context"
	"fmt"

	"github.com/dettarune/goTokoo/internal/entity"
	"github.com/dettarune/goTokoo/internal/model"
	"github.com/dettarune/goTokoo/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserUseCaseDeps struct {
	DB             *gorm.DB
	Log            *logrus.Logger
	Validate       *validator.Validate
	UserRepository *repository.UserRepository
}

func NewUserUseCase(dependencies *UserUseCaseDeps) *UserUseCaseDeps {
	return &UserUseCaseDeps{
		DB:       dependencies.DB,
		Log:      dependencies.Log,
		Validate: dependencies.Validate,
	}
}

func (c *UserUseCaseDeps) Register(ctx context.Context, request *model.RegsisterRequest) (*model.RegisterResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	// validate request
	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body: %+v", err)
		return nil, err
	}

	// check duplicate username
	total, err := c.UserRepository.CountByUsername(tx, request.Username)
	if err != nil {
		c.Log.Warnf("Failed count user from database : %+v", err)
		return nil, err
	}
	if total > 0 {
		return nil, fmt.Errorf("username already exists")
	}

	// hash password
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrypt hash : %+v", err)
		return nil, err
	}

	user := &entity.User{
		Username: request.Username,
		Password: string(password),
		FullName: request.FullName,
		Email:    request.Email,
	}

	// insert user
	if err := c.UserRepository.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, err
	}

	// commit transaction
	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, err
	}

	return &model.RegisterResponse{
		FullName:  user.FullName,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Unix(),
	}, nil
}
