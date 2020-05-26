package datastore

import (
	"context"
	"log"
	"time"

	"github.com/claudiocleberson/shippy-service-vessel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dabaseName     = "shippy"
	collectionName = "vessels"
)

var (
	mongoCollection *mongo.Collection
	retry           int
)

type MongoClient interface {
	FindAvailable(context.Context, *models.Specification) (*models.Vessel, error)
	Create(context.Context, *models.Vessel) error
}

func NewMongoClient(uri string) MongoClient {
	connectMongoCluster(uri)

	return &mongoClient{}
}

func connectMongoCluster(uri string) {

	log.Println("Connecting mongo cluster...")

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {

		if retry >= 3 {
			panic(err)
		}
		retry = retry + 1
		time.Sleep(time.Second * 2)
		connectMongoCluster(uri)
	}

	mongoCollection = client.Database(dabaseName).Collection(collectionName)

	log.Println("Mongo cluster connected...")
}

type mongoClient struct{}

func (c *mongoClient) FindAvailable(ctx context.Context, spec *models.Specification) (*models.Vessel, error) {

	filter := bson.D{{
		"maxweight", bson.D{{"$gt", spec.MaxWeight}},
	}}

	vessel := models.Vessel{}
	if err := mongoCollection.FindOne(ctx, filter).Decode(&vessel); err != nil {
		return nil, err
	}
	return &vessel, nil
}

func (c *mongoClient) Create(ctx context.Context, vessel *models.Vessel) error {
	_, err := mongoCollection.InsertOne(ctx, vessel)
	return err
}
