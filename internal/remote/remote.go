package remote

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/alimgiray/getir/adapter/mongo"
	"github.com/alimgiray/getir/internal/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	wrongParametersErr = errors.New("wrong parameters")
)

type RemoteHandler struct {
	DB *mongo.Mongo
}

func NewRemoteHandler() *RemoteHandler {
	return &RemoteHandler{
		DB: mongo.NewMongo(),
	}
}

func (h *RemoteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		h.handlePost(w, r)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *RemoteHandler) handlePost(w http.ResponseWriter, r *http.Request) {
	var request Request
	var response *Response

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		utils.ErrJSON(w, wrongParametersErr, http.StatusBadRequest)
		return
	}

	if request.StartDate == "" || request.EndDate == "" || request.MaxCount < request.MinCount {
		utils.ErrJSON(w, wrongParametersErr, http.StatusBadRequest)
		return
	}

	filter := bson.D{
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	option := bson.D{{Key: "_id", Value: 0}}
	cursor, err := h.DB.Col.Find(ctx, filter, options.Find().SetProjection(option))
	if err != nil {
		response = &Response{
			Code:    1,
			Message: "Failure - Query Failed",
		}
		utils.JSON(w, response, http.StatusInternalServerError)
		return
	}

	var results []*Record
	if err = cursor.All(ctx, &results); err != nil {
		response = &Response{
			Code:    1,
			Message: "Failure - Not Found Any Records",
		}
		utils.JSON(w, response, http.StatusNotFound)
		return
	}

	response = &Response{
		Code:    0,
		Message: "Success",
		Records: results,
	}

	utils.JSON(w, response, http.StatusOK)
}
