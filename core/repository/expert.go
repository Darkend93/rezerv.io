package repository

import (
	"github.com/google/uuid"
	"rezerv.io/core/domain/entity"
)

type Expert interface {
	GetByID(ID uuid.UUID) (*entity.Expert, error)
	GetAll() ([]*entity.Expert, error)
	Create(expert *entity.Expert) (*entity.Expert, error)
	Update(expert *entity.Expert) error
	Delete(expert *entity.Expert) error
}
