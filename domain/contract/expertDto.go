package contract

import (
	"github.com/google/uuid"
)

type ExpertDto struct {
	FirstName string
	LastName  string
	Services  []ServiceDto
}

type ServiceDto struct {
	ID      uuid.UUID
	Expert  string
	Price   float32
	Product string
}
