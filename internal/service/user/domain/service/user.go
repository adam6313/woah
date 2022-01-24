package service

import (
	"context"
	"woah/internal/service/user/domain/model/aggregate"

	"github.com/tyr-tech-team/hawk/status"
)

func (s *userService) CheckUserData(ctx context.Context, user *aggregate.User) error {

	var (
		errDesc = []string{}
	)

	// 檢查ID
	if user.ID == "" {
		errDesc = append(errDesc, "ID不可為空")

	}

	// 檢查姓名
	if user.Name == "" {
		errDesc = append(errDesc, "姓名不可為空")
	}

	if len(errDesc) > 0 {
		return status.InvalidParameter.WithDetail(errDesc...).Err()
	}

	return nil
}
