package redis

import (
	"context"
	"time"

	"github.com/artsadert/ShortLinker/internal/domain/entities"
	"github.com/artsadert/ShortLinker/internal/domain/repositories"
	"github.com/redis/go-redis/v9"
)

type RedisLinkRepository struct {
	db *redis.Client
}

func NewRedisLinkReepository(db *redis.Client) repositories.LinkRepository {
	return &RedisLinkRepository{db: db}
}

func (repo *RedisLinkRepository) Create(link *entities.Link) (string, error) {
	ctx := context.Background()

	short := link.Shorter()
	err := repo.db.Set(ctx, short, link.GetData(), 30*time.Second).Err()

	return short, err
}

func (repo *RedisLinkRepository) Get(link *entities.Link) (string, error) {
	ctx := context.Background()

	val, err := repo.db.Get(ctx, link.GetData()).Result()
	return val, err
}
