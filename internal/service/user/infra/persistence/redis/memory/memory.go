package memory

import (
	"context"
	"time"

	"github.com/tyr-tech-team/hawk/status"
)

// Get -
func (r *repo) Get(ctx context.Context, key string) (string, error) {
	if _, err := r.client.Do(ctx, "SELECT", r.db).Result(); err != nil {
		r.log(ctx).Errorf("redis connect failed: %s", err.Error())

		return "", status.ConnectFailed.WithDetail([]string{"redis connect failed"}...).Err()
	}

	return r.client.Get(ctx, key).Result()
}

// Delete -
func (r *repo) Delete(ctx context.Context, key ...string) (int64, error) {
	if _, err := r.client.Do(ctx, "SELECT", r.db).Result(); err != nil {
		r.log(ctx).Errorf("redis connect failed: %s", err.Error())
		return 0, status.ConnectFailed.WithDetail([]string{"redis connect failed"}...).Err()
	}

	count, err := r.client.Del(ctx, key...).Result()
	if err != nil {
		r.log(ctx).Errorf("redis delete failed: %s", err.Error())
		return 0, status.DeletedFailed.Err()
	}

	return count, nil
}

// SetNX -
func (r *repo) SetNX(ctx context.Context, key string, value interface{}, expire time.Duration) (bool, error) {
	if _, err := r.client.Do(ctx, "SELECT", r.db).Result(); err != nil {
		r.log(ctx).Errorf("redis connect failed: %s", err.Error())
		return false, status.ConnectFailed.WithDetail([]string{"redis connect failed"}...).Err()
	}
	return r.client.SetNX(ctx, key, value, expire).Result()
}
