package SpendsService

import (
	"log"

	"github.com/google/uuid"
)

type SpendsService interface {
	GetSpend(id uuid.UUID) (SpendResponseDTO, error)
	GetAllSpends() ([]SpendResponseDTO, error)
	CreateSpend(spend SpendRequestDTO) (SpendResponseDTO, error)
	UpdateSpend(spend SpendRequestDTO) (SpendResponseDTO, error)
	DeleteSpend(id uuid.UUID) error
}

type spendsService struct {
	repo SpendsRepository
}

func NewSpendsService(r SpendsRepository) SpendsService {
	return &spendsService{repo: r}
}

func (s *spendsService) GetSpend(id uuid.UUID) (SpendResponseDTO, error) {
	rawSpend, err := s.repo.GetSpend(id)
	if err != nil {
		return SpendResponseDTO{}, err
	}

	spend := MapSpendToResponseDTO(rawSpend)
	return spend, nil

}

func (s *spendsService) GetAllSpends() ([]SpendResponseDTO, error) {
	spends := []SpendResponseDTO{}

	rawSpends, err := s.repo.GetAllSpends()

	if err != nil {
		return []SpendResponseDTO{}, err
	}

	for _, v := range rawSpends {
		spend := MapSpendToResponseDTO(v)
		spends = append(spends, spend)
	}

	return spends, nil
}

func (s *spendsService) CreateSpend(spend SpendRequestDTO) (SpendResponseDTO, error) {
	log.Println("SERVICE, create:  request:", spend)

	rawSpend, err := MappingSpend(spend)
	if err != nil {
		return SpendResponseDTO{}, err
	}

	log.Println("SERVICE, create:  spend:", rawSpend)
	rawSpend, err = s.repo.CreateSpend(rawSpend)

	if err != nil {
		return SpendResponseDTO{}, err
	}

	newSpend := MapSpendToResponseDTO(rawSpend)
	log.Println("SERVICE, create: Responce:", rawSpend)
	return newSpend, nil
}

func (s *spendsService) UpdateSpend(spend SpendRequestDTO) (SpendResponseDTO, error) {
	log.Println("SERVICE, UpdateSpend: Responce:", spend)

	rawSpend, err := MappingSpend(spend)
	log.Println("SERVICE, UpdateSpend: spend:", spend)

	rawSpend, err = s.repo.UpdateSpend(rawSpend)
	if err != nil {
		return SpendResponseDTO{}, err
	}
	log.Println("SERVICE, UpdateSpend: res:", rawSpend)

	return MapSpendToResponseDTO(rawSpend), nil
}

func (s *spendsService) DeleteSpend(id uuid.UUID) error {

	log.Println("SRV DELETE:", id)

	return s.repo.DeleteSpend(id)
}
