package repository

import (
	"delegacia-facil/domain"

	"gorm.io/gorm"
)

type DelegaciaRepository interface {
	GetAll() ([]*domain.Delegacia, error)
	GetByHour(horario24h bool) ([]*domain.Delegacia, error)
}

type delegaciaRepository struct {
	db *gorm.DB
}

func NewdelegaciaRepository(db *gorm.DB) DelegaciaRepository {
	return &delegaciaRepository{db: db}
}

func (r *delegaciaRepository) GetByHour(horario24h bool) ([]*domain.Delegacia, error) {
	var delegacias []*domain.Delegacia

	err := r.db.Where("horario24h = ?", horario24h).Find(&delegacias).Error
	if err != nil {
		return nil, err
	}

	return delegacias, nil
}

func (r *delegaciaRepository) GetAll() ([]*domain.Delegacia, error) {
	var delegacias []*domain.Delegacia
	if err := r.db.Find(&delegacias).Error; err != nil {
		return nil, err
	}
	return delegacias, nil
}
