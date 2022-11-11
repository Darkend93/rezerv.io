package entity

import (
	"github.com/google/uuid"
)

type Expert struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Services  []Service
}
