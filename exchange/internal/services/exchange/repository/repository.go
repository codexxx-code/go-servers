package repository

import (
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"

	"exchange/internal/services/exchange/model"
	"pkg/cache"
	"pkg/database/geolite"
	"pkg/sql"
)

type ExchangeRepository struct {
	visibilityCache *cache.ItemCache[string, model.PublisherVisibility]
	pgsql           sql.SQL
	mongo           *mongo.Database
	geoLite         geolite.GeoLite
	redisADMs       *redis.Client
	redisUnusedBids *redis.Client
}

func NewExchangeRepository(
	pgsql sql.SQL,
	mongo *mongo.Database,
	geoLite geolite.GeoLite,
	redisADMs *redis.Client,
	redisUnusedBids *redis.Client,
) *ExchangeRepository {
	return &ExchangeRepository{
		visibilityCache: cache.NewItemCache[string, model.PublisherVisibility](5 * time.Second),
		pgsql:           pgsql,
		mongo:           mongo,
		geoLite:         geoLite,
		redisADMs:       redisADMs,
		redisUnusedBids: redisUnusedBids,
	}
}
