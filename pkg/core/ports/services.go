package ports

import (
	"github.com/matiasresca/xmen-meli/pkg/core/domain"
)

type HumanService interface {
	CheckMutant(dna []string) (bool, error)
	GetStats() (*domain.Stats, error)
}
