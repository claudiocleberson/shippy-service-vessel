package models

import (
	pb "github.com/claudiocleberson/shippy-service-vessel/proto/vessel"
)

type Specification struct {
	Capacity  int32 `json:"capacity"`
	MaxWeight int32 `json:"max_weight"`
}

func MarshalSpecification(spec *pb.Specification) *Specification {
	return &Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}

func UnmarshalSpecification(spec *Specification) *pb.Specification {
	return &pb.Specification{
		Capacity:  spec.Capacity,
		MaxWeight: spec.MaxWeight,
	}
}
