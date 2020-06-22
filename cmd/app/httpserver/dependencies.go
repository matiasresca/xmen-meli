package httpserver

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/matiasresca/xmen-meli/pkg/core/ports"
	"github.com/matiasresca/xmen-meli/pkg/core/services/humansrv"
	"github.com/matiasresca/xmen-meli/pkg/infrastructure/handler/humanhdl"
	"github.com/matiasresca/xmen-meli/pkg/infrastructure/repositories/humanrepo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	//Configuracion por defecto.-
	host := "localhost"
	port := 27017
	dbName := "xmendb"
	//Valores obtenidos de las variables de entorno.-
	if hostEnv := os.Getenv("MONGO_HOST"); hostEnv != "" {
		host = hostEnv
	}
	if portEnv := os.Getenv("MONGO_PORT"); portEnv != "" {
		portInt, err := strconv.Atoi(portEnv)
		if err != nil {
			panic(err)
		}
		port = portInt
	}
	if dbNameEnv := os.Getenv("MONGO_DB"); dbNameEnv != "" {
		dbName = dbNameEnv
	}
	//Obtengo conexion.-
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
