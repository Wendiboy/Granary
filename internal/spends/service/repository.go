package SpendsService

import (
	"fmt"
	"log"

	"github.com/google/uuid"
)

type SpendsRepository interface {
	GetSpend(id uuid.UUID) (RawSpend, error)
	GetAllSpends() ([]RawSpend, error)
	CreateSpend(rawSpend RawSpend) (RawSpend, error)
	UpdateSpend(rawSpend RawSpend) (RawSpend, error)
	DeleteSpend(id uuid.UUID) error
}

type spendsRepository struct {
}

func NewSpendsRepository() SpendsRepository {
	return &spendsRepository{}
}

var DB []RawSpend

func (r *spendsRepository) GetSpend(id uuid.UUID) (RawSpend, error) {
	log.SetPrefix("REPO: GetSpend")

	for _, v := range DB {
		log.Println(v)
		if v.Id == id {
			return v, nil
		}
	}

	err := fmt.Errorf("spend %d was not found", id)

	return RawSpend{}, err
}

func (r *spendsRepository) GetAllSpends() ([]RawSpend, error) {
	log.SetPrefix("REPO: GetAllSpends")

	rawSpends := make([]RawSpend, 0, 10)

	for _, v := range DB {
		rawSpends = append(rawSpends, v)
	}

	log.Println(rawSpends)
	return rawSpends, nil
}

func (r *spendsRepository) CreateSpend(rawSpend RawSpend) (RawSpend, error) {
	log.SetPrefix("REPO: create")
	DB = append(DB, rawSpend)
	log.Println(rawSpend)
	return rawSpend, nil
}

func (r *spendsRepository) UpdateSpend(rawSpend RawSpend) (RawSpend, error) {
	log.SetPrefix("REPO: Update")
	for i, v := range DB {
		if v.Id == rawSpend.Id {
			DB[i] = rawSpend
		}
	}
	log.Println(rawSpend)

	return rawSpend, nil
}

func (r *spendsRepository) DeleteSpend(id uuid.UUID) error {
	for i, v := range DB {
		if id == v.Id {
			DB = append(DB[:i], DB[i+1:]...)
			return nil
		}
	}
	return nil
}
