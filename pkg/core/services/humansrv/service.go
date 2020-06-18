package humansrv

import "github.com/matiasresca/xmen-meli/pkg/core/ports"

type service struct {
	repo ports.HumanRepository
}

func NewService(repo ports.HumanRepository) ports.HumanService {
	return &service{
		repo: repo,
	}
}

func (s service) IsMutant(dna []string) bool {
	panic("implement me")
}
