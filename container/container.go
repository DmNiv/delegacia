package container

import (
	"test/adapter/http/handler"
	"test/adapter/repository"
	"test/core/usecase"
	"test/postgres"

	"gorm.io/gorm"
)

type Container struct {
	DB               *gorm.DB
	UserRepo         repository.UserRepository
	UserUseCase      *usecase.UserUseCase
	UserHandler      *handler.UserHandler
	DelegaciaRepo    repository.DelegaciaRepository
	DelegaciaUseCase *usecase.DelegaciaUseCase
	DelegaciaHandler *handler.DelegaciaHandler
}
 
func NewContainer() *Container {
	// Inicializa o banco de dados
	db := postgres.InitDB()

	// Criação dos repositórios
	userRepo := repository.NewUserRepository(db)
	delegaciasRepo := repository.NewdelegaciaRepository(db)

	// Criação dos casos de uso
	userUseCase := usecase.NewUserUseCase(userRepo)
	delegaciaUseCase := usecase.NewDelegaciaUseCase(delegaciasRepo)

	// Criação dos handlers
	userHandler := handler.NewUserHandler(userUseCase)
	delegaciaHandler := handler.NewDelegaciaHandler(delegaciaUseCase)

	return &Container{
		DB:          db,

		UserRepo:    userRepo,
		UserUseCase: userUseCase,
		UserHandler: userHandler,

		DelegaciaRepo: delegaciasRepo,
		DelegaciaUseCase: delegaciaUseCase,
		DelegaciaHandler: delegaciaHandler,
	}
}
