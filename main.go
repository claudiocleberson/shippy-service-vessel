package main

import (
	pb "github.com/claudiocleberson/shippy-service-vessel/proto/vessel"
	"github.com/claudiocleberson/shippy-service-vessel/repository"
	"github.com/claudiocleberson/shippy-service-vessel/service"
	"github.com/micro/go-micro"
)

func main() {
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},
		&pb.Vessel{Id: "vessel002", Name: "Boaty McBoatface II", MaxWeight: 240000, Capacity: 600},
		&pb.Vessel{Id: "vessel003", Name: "Boaty McBoatface III", MaxWeight: 300000, Capacity: 700},
		&pb.Vessel{Id: "vessel004", Name: "Boaty McBoatface III", MaxWeight: 500000, Capacity: 1000},
	}

	repo := repository.NewVesselRepository(vessels)
	vesselService := service.NewVesselService(repo)

	srv := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)
	srv.Init()

	pb.RegisterVesselServiceHandler(srv.Server(), vesselService)

	if err := srv.Run(); err != nil {
		panic(err)
	}

}
