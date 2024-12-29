package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"exchange/internal/services/event/repository/eventDDL"
	"exchange/internal/services/exchange/model"
	"pkg/errors"
)

//

func (r *ExchangeRepository) GetPublisherVisibility(ctx context.Context, publisherID string) (visibility model.PublisherVisibility, err error) {

	// Получаем из кэша значение видимости по паблишеру
	var ok bool
	visibility, ok = r.visibilityCache.Get(publisherID)
	if !ok { // Если данных нет или они устарели

		// Получаем данные из базы данных
		if err := r.mongo.Collection(eventDDL.Collection).FindOne(
			ctx,
			// Filter
			bson.M{
				"_id": publisherID,
			},
		).Decode(&visibility); err != nil {
			switch {
			case errors.Is(err, mongo.ErrNoDocuments):
				return model.PublisherVisibility{
					PublisherID: publisherID,
					Loads:       0,
					Views:       0,
				}, nil
			case errors.IsContextError(err):
				return visibility, errors.Timeout.Wrap(err)
			default:
				return visibility, errors.InternalServer.Wrap(err)
			}
		}
	}

	return visibility, nil
}
