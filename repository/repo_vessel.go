package repository

import (
	"context"

	"github.com/claudiocleberson/shippy-service-vessel/datastore"
	"github.com/claudiocleberson/shippy-service-vessel/models"
)

type VesselRepository interface {
	FindAvailable(context.Context, *models.Specification) (*models.Vessel, error)
	Create(context.Context, *models.Vessel) error
}

type vesselRepository struct {
	mongoClient datastore.MongoClient
}

func NewVesselRepository(cli datastore.MongoClient) VesselRepository {
	return &vesselRepository{
		mongoClient: cli,
	}
}

func (repo *vesselRepository) FindAvailable(ctx context.Context, spec *models.Specification) (*models.Vessel, error) {

	vessel, err := repo.mongoClient.FindAvailable(ctx, spec)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repo *vesselRepository) Create(ctx context.Context, vessel *models.Vessel) error {
	err := repo.mongoClient.Create(ctx, vessel)
	return err
}
