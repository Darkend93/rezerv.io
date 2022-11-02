package service

import (
	"github.com/google/uuid"
	"rezerv.io/domain/contract"
	"rezerv.io/domain/entity"
	"rezerv.io/usecase/repository"
)

type expertService struct {
	Repository repository.ExpertRepository
}

func NewExpertService(e repository.ExpertRepository) ExpertService {
	return &expertService{Repository: e}
}

type ExpertService interface {
	Register(expert *contract.ExpertDto) error
	AddService(val *contract.ExpertServiceDto) error
	Disable() error
}

func (es *expertService) Register(expert *contract.ExpertDto) error {
	expertEntity := entity.Expert{
		ID:        uuid.New(),
		FirstName: expert.FirstName,
		LastName:  expert.LastName,
		Services:  nil,
	}

	_, err := es.Repository.Create(&expertEntity)
	if err != nil {
		return err
	}

	return nil
}

func (es *expertService) AddService(val *contract.ExpertServiceDto) error {
	expert, err := es.Repository.GetByID(val.ExpertID)
	if err != nil {
		return err
	}

	return nil
}

func (es *expertService) Disable() error {
	return nil
}
