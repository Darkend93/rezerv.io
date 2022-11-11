package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"rezerv.io/domain/contract"
	"rezerv.io/usecase/repository"
)

type rezerv struct {
	logger     *log.Logger
	repository repository.Rezerv
}

type Rezerv interface {
	MakeRezerv(ctx context.Context, rezerv *contract.Rezervation) error
	CompleteRezerv(ctx context.Context, rezervId uuid.UUID) error
}

func NewRezerv(r repository.Rezerv, l *log.Logger) Rezerv {
	return &rezerv{
		repository: r,
		logger:     l,
	}
}

func (r *rezerv) MakeRezerv(ctx context.Context, rezerv *contract.Rezervation) error {
	if _, err := r.repository.Create(rezerv.ToEntity()); err != nil {
		r.logger.Fatalln("Unable to create a new Rezerv in DB error: {1}", err)
		return err
	}

	return nil
}

func (r *rezerv) CompleteRezerv(ctx context.Context, rezervId uuid.UUID) error {
	rzrv, err := r.repository.GetByID(rezervId)
	if err != nil {
		r.logger.Fatalln("Rezerv not found error: {1}", err)
		return err
	}

	rzrv.IsVisited = true

	if err = r.repository.Update(rzrv); err != nil {
		r.logger.Fatalln("Failed to update record error: {1}", err)
		return err
	}

	return nil
}
