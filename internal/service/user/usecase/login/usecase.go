package login

import (
	"context"
	"time"
	"woah/internal/service/user/domain/model/value_object"
	"woah/internal/service/user/domain/repository"
	"woah/internal/service/user/domain/service"
)

// loginUsecase -
type loginUsecase struct {
	AggregateID string
	userRepo    repository.UserRepository
	userRedis   repository.RedisRepository
	userService service.UserService
}

// NewUseCase -
func NewUseCase(userRepo repository.UserRepository, userRedis repository.RedisRepository, userService service.UserService) LoginUserUsecase {
	return &loginUsecase{
		userRepo:    userRepo,
		userRedis:   userRedis,
		userService: userService,
	}
}

// Login - 登入
func (l *loginUsecase) Login(ctx context.Context, command *UserLoginCmd) (*UserLogined, error) {

	userToken := &value_object.UserToken{
		ID:   l.AggregateID,
		Name: command.Name,
	}
	userToken.DefaultClaims()

	// 取得User
	user, err := l.userRepo.GetUser(ctx, userToken.ID)
	if err != nil {
		return nil, err
	}

	// 產生Token
	token, err := l.userService.GenerateToken(ctx, userToken.Claims)
	if err != nil {
		return nil, err
	}

	// 暫存Token
	if _, err := l.userRedis.SetNX(ctx, user.ID, token, time.Hour*24); err != nil {
		return nil, err
	}

	return &UserLogined{
		// Token -
		Token: token,
	}, err
}
