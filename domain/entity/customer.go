package entity

import "github.com/google/uuid"

type Customer struct {
	ID          uuid.UUID
	FirstName   string
	LastName    string
	MobilePhone string
}
