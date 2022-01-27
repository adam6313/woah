package login

import "context"

type LoginUserUsecase interface {
	// Login - 使用者登入
	Login(ctx context.Context, command *UserLoginCmd) (*UserLogined, error)
}
