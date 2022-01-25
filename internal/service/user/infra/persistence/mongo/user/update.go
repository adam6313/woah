package user

import (
	"context"
	"woah/internal/service/user/domain/model/aggregate"

	"github.com/tyr-tech-team/hawk/status"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateUser - 更新使用者資料
func (r *repo) UpdateUser(ctx context.Context, user *aggregate.User) error {

	c := r.db.Database(r.databaseName).Collection(C)

	// set filter
	filter := bson.M{
		"Id": user.ID,
	}

	// set updater
	updater := bson.M{
		"$set": bson.M{
			"name": user.Name,
		},
	}

	if _, err := c.UpdateOne(ctx, filter, updater); err != nil {
		r.log(ctx).Error(err)
		st := status.UpdatedFailed.WithDetail([]string{"更新使用者資料失敗"}...).Err()
		return st
	}

	return nil
}
