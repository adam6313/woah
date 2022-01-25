package user

import (
	"context"
	"woah/internal/service/user/domain/model/aggregate"
	"woah/internal/service/user/infra/persistence/mongo/po"

	"github.com/tyr-tech-team/hawk/status"
)

// CreateUser - 建立使用者
func (r *repo) CreateUser(ctx context.Context, user *aggregate.User) error {

	c := r.db.Database(r.databaseName).Collection(C)

	// po
	userPo := po.ConverToUserPo(user)

	if _, err := c.InsertOne(ctx, user); err != nil {
		r.log(ctx).Errorf("create user failed with args [%+v]: %s", userPo, err.Error())
		return status.CreatedFailed.Err()
	}

	return nil

}
