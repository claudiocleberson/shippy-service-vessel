package models

import (
	pb "github.com/claudiocleberson/shippy-service-vessel/proto/vessel"
)

type Vessel struct {
	ID        string `json:"id"`
	Capacity  int32  `json:"capacity"`
	Name      string `json:"name"`
	Available bool   `json:"available"`
	OwnerID   string `json:"owner_id"`
	MaxWeight int32  `json:"max_weight"`
}

func MarshalVessel(vessel *pb.Vessel) *Vessel {
	return &Vessel{
		ID:        vessel.Id,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerID:   vessel.OwnerId,
	}
}

func UnmarshalVessel(vessel *Vessel) *pb.Vessel {
	return &pb.Vessel{
		Id:        vessel.ID,
		Capacity:  vessel.Capacity,
		MaxWeight: vessel.MaxWeight,
		Name:      vessel.Name,
		Available: vessel.Available,
		OwnerId:   vessel.OwnerID,
	}
}
