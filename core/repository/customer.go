package repository

import (
	"github.com/google/uuid"
	"rezerv.io/core/domain/entity"
)

type Customer interface {
	GetByID(ID uuid.UUID) (*entity.Customer, error)
	GetAll() ([]*entity.Customer, error)
	Create(customer *entity.Customer) (*entity.Customer, error)
	Update(customer *entity.Customer) error
	Delete(customer *entity.Customer) error
}
