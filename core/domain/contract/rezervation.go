package contract

import (
	"time"

	"github.com/google/uuid"
	"rezerv.io/core/domain/entity"
)

type Rezervation struct {
	Customer uuid.UUID
	Date     time.Time
	Service  uuid.UUID
}

func (r *Rezervation) ToEntity() *entity.Rezerv {
	return &entity.Rezerv{
		ID:         uuid.New(),
		CustomerID: r.Customer,
		Date:       r.Date,
		IsVisited:  false,
		Service:    r.Service,
	}
}
