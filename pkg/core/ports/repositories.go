package ports

import "github.com/matiasresca/xmen-meli/pkg/core/domain"

type HumanRepository interface {
	Save(human domain.Human) error
	GetByDna(dna []string) (*domain.Human, error)
	GetAll() ([]domain.Human, error)
}
