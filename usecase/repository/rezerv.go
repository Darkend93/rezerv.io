package repository

import (
	"github.com/google/uuid"
	"rezerv.io/domain/entity"
)

type Rezerv interface {
	GetByID(ID uuid.UUID) (*entity.Rezerv, error)
	GetAll() ([]*entity.Rezerv, error)
	Create(rezerv *entity.Rezerv) (*entity.Rezerv, error)
	Update(rezerv *entity.Rezerv) error
	Delete(rezerv *entity.Rezerv) error
}
