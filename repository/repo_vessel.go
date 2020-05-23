package repository

import (
	"errors"

	pb "github.com/claudiocleberson/shippy-service-vessel/proto/vessel"
)

type VesselRepository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type vesselRepository struct {
	vessels []*pb.Vessel
}

func NewVesselRepository(vessels []*pb.Vessel) VesselRepository {
	return &vesselRepository{
		vessels: vessels,
	}
}

func init() {

}

func (repo *vesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {

	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}
