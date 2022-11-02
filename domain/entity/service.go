package entity

import "github.com/google/uuid"

type Service struct {
	ID      uuid.UUID
	Expert  Expert
	Price   float32
	Product Product
}
