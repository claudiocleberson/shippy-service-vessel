package main

import (
	"os"

	"github.com/claudiocleberson/shippy-service-vessel/datastore"
	"github.com/claudiocleberson/shippy-service-vessel/handlers"
	pb "github.com/claudiocleberson/shippy-service-vessel/proto/vessel"
	"github.com/claudiocleberson/shippy-service-vessel/repository"
	"github.com/micro/go-micro"
)

const (
	port        = ":50052"
	dbHost      = "DB_HOST"
	defaultHost = "mongodb://localhost:27017"
)

func main() {

	mongoUri := os.Getenv(dbHost)
	if mongoUri == "" {
		mongoUri = defaultHost
	}

	// vessels := []*pb.Vessel{
	// 	{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
	// 	{Id: "vessel002", Name: "Boaty McBoatface II", MaxWeight: 240000, Capacity: 600},
	// 	{Id: "vessel003", Name: "Boaty McBoatface III", MaxWeight: 300000, Capacity: 700},
	// 	{Id: "vessel004", Name: "Boaty McBoatface III", MaxWeight: 500000, Capacity: 1000},
	// }

	dbClient := datastore.NewMongoClient(mongoUri)

	repo := repository.NewVesselRepository(dbClient)

	vesselService := handlers.NewVesselServiceHandler(repo)

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), vesselService)

	if err := srv.Run(); err != nil {
		panic(err)
	}

}
