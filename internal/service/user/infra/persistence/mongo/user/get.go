package user

import (
	"context"
	"woah/internal/service/user/domain/model/aggregate"

	"github.com/tyr-tech-team/hawk/status"
	"go.mongodb.org/mongo-driver/bson"
)

// GetUser - 取得使用者
func (r *repo) GetUser(ctx context.Context, id string) (*aggregate.User, error) {

	c := r.db.Database(r.databaseName).Collection(C)

	filter := bson.M{
		"Id": id,
	}

	cure := c.FindOne(ctx, filter)

	user := new(aggregate.User)

	if err := cure.Decode(&user); err != nil {
		r.log(ctx).Error(err)

		st := status.NotFound.WithDetail([]string{"取得使用者資料失敗"}...).Err()
		return nil, st

	}

	return user, nil
}
