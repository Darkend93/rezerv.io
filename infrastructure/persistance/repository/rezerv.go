package repository

import (
	"log"

	"github.com/google/uuid"
	"rezerv.io/core/domain/entity"
	ir "rezerv.io/core/repository"

	"gorm.io/gorm"
)

type rezerv struct {
	logger *log.Logger
	db     *gorm.DB
}

func NewRepository(l *log.Logger, d *gorm.DB) ir.Rezerv {
	return &rezerv{
		logger: l,
		db:     d,
	}
}

func (r *rezerv) GetByID(ID uuid.UUID) (*entity.Rezerv, error) {
	return nil, nil
}

func (r *rezerv) GetAll() ([]*entity.Rezerv, error) {
	return nil, nil
}

func (r *rezerv) Create(rezerv *entity.Rezerv) (*entity.Rezerv, error) {
	return nil, nil
}
func (r *rezerv) Update(rezerv *entity.Rezerv) error {
	return nil
}
func (r *rezerv) Delete(rezerv *entity.Rezerv) error {
	return nil
}
