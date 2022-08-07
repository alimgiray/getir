package remote

import (
	"context"
	"time"

	"github.com/alimgiray/getir/adapter/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	SUCCESS int = iota
	FAILURE
)

type RemoteService struct {
	db *mongo.Mongo
}

func NewRemoteService() *RemoteService {
	return &RemoteService{db: mongo.NewMongo()}
}

func (s *RemoteService) FindRecord(request Request) (int, string, []*Record) {
	filter := s.createFilterFromRequest(request)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := s.db.Col.Find(ctx, filter, options.Find().SetProjection(bson.D{{Key: "_id", Value: 0}}))
	if err != nil {
		return FAILURE, "Failure - Query Failed", nil
	}

	var results []*Record
	if err = cursor.All(ctx, &results); err != nil {
		return FAILURE, "Failure - Not Found Any Records", nil
	}

	return SUCCESS, "Success", results
}

func (s *RemoteService) createFilterFromRequest(request Request) bson.D {
	return bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{
					{Key: "createdAt", Value: bson.D{{Key: "$gt", Value: request.StartDate}}},
				},
				bson.D{
					{Key: "createdAt", Value: bson.D{{Key: "$lt", Value: request.EndDate}}},
				},
				bson.D{
					{Key: "totalCount", Value: bson.D{{Key: "$gt", Value: request.MinCount}}},
				},
				bson.D{
					{Key: "totalCount", Value: bson.D{{Key: "$lt", Value: request.MaxCount}}},
				},
			},
		},
	}
}
