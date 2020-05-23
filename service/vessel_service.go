package service

import (
	"context"

	pb "github.com/claudiocleberson/shippy-service-vessel/proto/vessel"
	"github.com/claudiocleberson/shippy-service-vessel/repository"
)

type VesselService interface {
	FindAvailable(context.Context, *pb.Specification, *pb.Response) error
}

type vesselService struct {
	repo repository.VesselRepository
}

func NewVesselService(repo repository.VesselRepository) VesselService {
	return &vesselService{
		repo: repo,
	}
}

func (s *vesselService) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	//Set the vessel as part o the response message type
	res.Vessel = vessel
	return nil
}
