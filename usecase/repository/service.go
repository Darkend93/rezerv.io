package repository

import (
	"github.com/google/uuid"
	"rezerv.io/domain/entity"
)

type Service interface {
	GetByID(ID uuid.UUID) (*entity.Service, error)
	GetAll() ([]*entity.Service, error)
	Create(service *entity.Service) (*entity.Service, error)
	Update(service *entity.Service) error
	Delete(service *entity.Service) error
}
