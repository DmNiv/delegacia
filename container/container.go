package container

import (
	"test/adapter/http/handler"
	"test/adapter/repository"
	"test/core/usecase"
	"test/postgres"

	"gorm.io/gorm"
)

type Container struct {
	DB          *gorm.DB
	UserRepo    repository.UserRepository
	UserUseCase *usecase.UserUseCase
	UserHandler *handler.UserHandler
}

func NewContainer() *Container {
	// Inicializa o banco de dados
	db := postgres.InitDB()

	// Criação dos repositórios
	userRepo := repository.NewUserRepository(db)

	// Criação dos casos de uso
	userUseCase := usecase.NewUserUseCase(userRepo)

	// Criação dos handlers
	userHandler := handler.NewUserHandler(userUseCase)

	return &Container{
		DB:          db,
		UserRepo:    userRepo,
		UserUseCase: userUseCase,
		UserHandler: userHandler,
	}
}
