package ports

import "github.com/matiasresca/xmen-meli/pkg/core/domain"

type HumanRepository interface {
	Save(human domain.Human) error
}
