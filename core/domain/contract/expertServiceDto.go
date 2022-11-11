package contract

import "github.com/google/uuid"

type ExpertServiceDto struct {
	ExpertID  uuid.UUID
	ProductID uuid.UUID
	Price     float32
}
