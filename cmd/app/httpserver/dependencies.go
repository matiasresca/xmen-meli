package httpserver

import (
	"context"
	"fmt"
	"github.com/matiasresca/xmen-meli/pkg/core/ports"
	"github.com/matiasresca/xmen-meli/pkg/core/services/humansrv"
	"github.com/matiasresca/xmen-meli/pkg/infrastructure/handler/humanhdl"
	"github.com/matiasresca/xmen-meli/pkg/infrastructure/repositories/humanrepo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	Dependencies *DependenciesDefinitions = &DependenciesDefinitions{}
)

type DependenciesDefinitions struct {
	HumanRepository ports.HumanRepository
	HumanService    ports.HumanService
	HumanHandler    *humanhdl.HTTPHandler
}

func (d *DependenciesDefinitions) Initialize() {
	d.HumanRepository = humanrepo.NewMongo(newMongoConection())
	d.HumanService = humansrv.NewService(d.HumanRepository)
	d.HumanHandler = humanhdl.NewHTTPHandler(d.HumanService)
}

func newMongoConection() *mongo.Database {
	host := "localhost"
	port := 27017
	dbName := "xmendb"

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connections
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(dbName)
}
