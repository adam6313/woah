package service

import (
	"context"
	"woah/internal/service/user/domain/model/aggregate"

	"github.com/dgrijalva/jwt-go"
	"github.com/tyr-tech-team/hawk/status"
)

var secretKey = []byte("tyr-woah-9527")

// GenerateToken -
func (s *userService) GenerateToken(ctx context.Context, data jwt.Claims) (string, error) {
	// new jwt
	t := jwt.NewWithClaims(jwt.SigningMethodHS512, data)

	// make jwt to string
	token, err := t.SignedString(secretKey)
	if err != nil {
		s.log(ctx).Errorf(err.Error())
		return "", status.TokenGenerationFailed.Err()
	}
	return token, nil
}

// CheckUserData - 檢查使用者資料
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

// SetUpdateUser - 設置使用者更新內容
// origin - 原始資料
// in - 欲更新的資料
func (s *userService) SetUpdateUser(ctx context.Context, origin, in *aggregate.User) error {

	// 更新名字
	if in.Name != "" && in.Name != origin.Name {
		origin.Name = in.Name
	}
	return nil
}
