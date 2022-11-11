package entity

import (
	"time"

	"github.com/google/uuid"
)

type Rezerv struct {
	ID         uuid.UUID
	CustomerID uuid.UUID
	Date       time.Time
	IsVisited  bool
	Service    uuid.UUID
}
