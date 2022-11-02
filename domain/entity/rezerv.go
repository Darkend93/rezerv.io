package entity

import (
	"time"

	"github.com/google/uuid"
)

type Rezerv struct {
	ID        uuid.UUID
	Customer  Customer
	Date      time.Time
	IsVisited bool
	Service   Service
}
