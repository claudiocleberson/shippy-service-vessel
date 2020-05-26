package handlers

import (
	"context"

	"github.com/claudiocleberson/shippy-service-vessel/models"
	pb "github.com/claudiocleberson/shippy-service-vessel/proto/vessel"
	"github.com/claudiocleberson/shippy-service-vessel/repository"
)

type VesselServiceHandler interface {
	FindAvailable(context.Context, *pb.Specification, *pb.Response) error
	Create(context.Context, *pb.Vessel, *pb.Response) error
}

type vesselServiceHandler struct {
	repo repository.VesselRepository
}

func NewVesselServiceHandler(repo repository.VesselRepository) VesselServiceHandler {
	return &vesselServiceHandler{
		repo: repo,
	}
}

func (s *vesselServiceHandler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	vessel, err := s.repo.FindAvailable(ctx, models.MarshalSpecification(req))
	if err != nil {

		//{Id: "vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},

		s.repo.Create(ctx, &models.Vessel{
			ID:        "vessel12020",
			Name:      "boaty killer",
			MaxWeight: 300000,
			Capacity:  10000,
		})

		s.repo.Create(ctx, &models.Vessel{
			ID:        "vessel12021",
			Name:      "boaty mcboatface",
			MaxWeight: 250000,
			Capacity:  10000,
		})
		s.repo.Create(ctx, &models.Vessel{
			ID:        "vessel12022",
			Name:      "boaty Panter",
			MaxWeight: 450000,
			Capacity:  10000,
		})
		s.repo.Create(ctx, &models.Vessel{
			ID:        "vessel12023",
			Name:      "boaty b",
			MaxWeight: 500000,
			Capacity:  10000,
		})

		return err
	}

	//Set the vessel as part o the response message type
	res.Vessel = models.UnmarshalVessel(vessel)
	return nil
}

func (s *vesselServiceHandler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {

	err := s.repo.Create(ctx, models.MarshalVessel(req))
	if err != nil {
		return err
	}

	res.Vessel = req

	return nil
}
