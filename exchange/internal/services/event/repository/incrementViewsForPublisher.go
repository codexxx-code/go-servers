package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"exchange/internal/services/event/repository/eventDDL"
	"pkg/errors"
)

func (r *EventRepository) IncrementViewsForPublisher(ctx context.Context, publisherID string) error {

	_, err := r.mongo.Collection(eventDDL.Collection).UpdateOne(ctx,
		// Filter
		bson.M{
			"_id": publisherID,
		},
		// Update
		bson.M{
			"$inc": bson.M{
				"views": 1,
			},
		},
	)
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	return nil
}
