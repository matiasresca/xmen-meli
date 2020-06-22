package humanrepo

import (
	"context"
	"log"
	"sync"

	"github.com/matiasresca/xmen-meli/pkg/core/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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

func (repo *mongoHumanRepository) Save(human domain.Human) error {
	_, err := repo.db.Collection("human").InsertOne(context.TODO(), human)
	if err != nil {
		return err
	}
	return nil
}

func (repo *mongoHumanRepository) GetByDna(dna []string) (*domain.Human, error) {
	var human domain.Human
	err := repo.db.Collection("human").FindOne(context.TODO(), bson.D{{"dna", dna}}).Decode(&human)
	if err != nil {
		return nil, err
	}
	return &human, nil
}

func (repo *mongoHumanRepository) GetAll() ([]domain.Human, error) {
	humans := []domain.Human{}
	//Busco todos los registros en la db.-
	ctx := context.Background()
	cursor, err := repo.db.Collection("human").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	//Por cada cursor seteo el modelo.-
	for cursor.Next(ctx) {
		var human domain.Human
		//Decode del registro al modelo.-
		if err := cursor.Decode(&human); err != nil {
			log.Fatal("cursor.Decode ERROR:", err)
		}
		//Agrego valores al slice.-
		humans = append(humans, human)
	}
	return humans, nil
}
