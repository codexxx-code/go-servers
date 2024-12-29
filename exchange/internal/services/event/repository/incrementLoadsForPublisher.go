package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"exchange/internal/services/event/repository/eventDDL"
	"pkg/errors"
	"pkg/pointer"
)

func (r *EventRepository) IncrementLoadsForPublisher(ctx context.Context, publisherID string) error {

	_, err := r.mongo.Collection(eventDDL.Collection).UpdateOne(ctx,
		// Filter
		bson.M{
			"_id": publisherID,
		},
		// Update
		bson.M{
			"$inc": bson.M{
				"loads": 1,
			},
		},
		&options.UpdateOptions{ //nolint:exhaustruct
			Upsert: pointer.Pointer(true),
		},
	)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
