package usecase

import (
	"delegacia-facil/adapter/repository"
	"delegacia-facil/domain"
)

type DelegaciaUseCase struct {
	delegaciaRepo repository.DelegaciaRepository
}

func NewDelegaciaUseCase(delegaciaRepo repository.DelegaciaRepository) *DelegaciaUseCase {
	return &DelegaciaUseCase{delegaciaRepo: delegaciaRepo}
}

func (dl *DelegaciaUseCase) ListaDelegacias() ([]*domain.Delegacia, error) {
	delegacias, err := dl.delegaciaRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return delegacias, nil
}

func (dl *DelegaciaUseCase) FiltraDelegacias(booleano bool) ([]*domain.Delegacia, error) {
	delegacias, err := dl.delegaciaRepo.GetByHour(booleano)
	if err != nil {
		return nil, err
	}

	return delegacias, nil
}
