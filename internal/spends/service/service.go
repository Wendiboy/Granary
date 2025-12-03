package SpendsService

import (
	"log"

	"github.com/google/uuid"
)

type SpendsService interface {
	GetSpend(id uuid.UUID) (Spend, error)
	GetAllSpends() ([]Spend, error)
	CreateSpend(spend Spend) (Spend, error)
	UpdateSpend(id uuid.UUID, spend Spend) (Spend, error)
	DeleteSpend(id uuid.UUID) error
}

type spendsService struct {
	repo SpendsRepository
}

func NewSpendsService(r SpendsRepository) SpendsService {
	return &spendsService{repo: r}
}

func (s *spendsService) GetSpend(id uuid.UUID) (Spend, error) {
	rawSpend, err := s.repo.GetSpend(id)
	if err != nil {
		return Spend{}, err
	}

	spend := MappingSpend(rawSpend)
	return spend, err

}

func (s *spendsService) GetAllSpends() ([]Spend, error) {
	spends := []Spend{}

	rawSpends, err := s.repo.GetAllSpends()

	if err != nil {
		return []Spend{}, err
	}

	for _, v := range rawSpends {
		spend := MappingSpend(v)
		spends = append(spends, spend)
	}

	return spends, nil
}

func (s *spendsService) CreateSpend(spend Spend) (Spend, error) {
	log.SetPrefix("SRV, create:")
	log.Println("reqSpend:", spend)

	id, _ := uuid.NewUUID()

	rawSpend, err := ReMappingSpend(id, spend)
	log.Println("to DB:", rawSpend)

	rawSpend, err = s.repo.CreateSpend(rawSpend)

	if err != nil {
		return Spend{}, err
	}

	newSpend := MappingSpend(rawSpend)

	return newSpend, nil
}

func (s *spendsService) UpdateSpend(id uuid.UUID, spend Spend) (Spend, error) {
	log.SetPrefix("SRV Update:")

	rawSpend, err := ReMappingSpend(id, spend)
	log.Println(rawSpend)

	rawSpend, err = s.repo.UpdateSpend(rawSpend)
	if err != nil {
		return Spend{}, err
	}

	spend = MappingSpend(rawSpend)

	return spend, nil
}

func (s *spendsService) DeleteSpend(id uuid.UUID) error {

	log.Println("SRV DELETE:", id)

	return s.repo.DeleteSpend(id)
}
