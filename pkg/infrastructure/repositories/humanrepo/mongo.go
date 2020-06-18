package humanrepo

import (
	"github.com/matiasresca/xmen-meli/pkg/core/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var (
	mongoOnce     sync.Once
	instanceMongo *mongoHumanRepository
)

type mongoHumanRepository struct {
	db *mongo.Database
}

func NewMongo(db *mongo.Database) *mongoHumanRepository {
	mongoOnce.Do(func() {
		instanceMongo = &mongoHumanRepository{db: db}
	})
	return instanceMongo
}

func (m mongoHumanRepository) Save(human domain.Human) error {
	panic("implement me")
}
