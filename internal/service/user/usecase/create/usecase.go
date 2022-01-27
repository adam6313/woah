package create

import (
	"context"
	"woah/internal/service/user/domain/model/aggregate"
	"woah/internal/service/user/domain/repository"
	"woah/internal/service/user/domain/service"
)

// CreateUsecase -
type createUsecase struct {
	userRepo    repository.UserRepository
	userService service.UserService
	AggregateID string
}

// NewUseCase -
func NewUseCase(userRepo repository.UserRepository, userService service.UserService) CreateUserUsecase {
	return &createUsecase{
		userRepo:    userRepo,
		userService: userService,
	}
}

// Create -
func (c *createUsecase) Create(ctx context.Context, command *CreateUser) (*UserCreated, error) {
	//aggregateID := ctx.Value("aggregateID").(string)

	user := &aggregate.User{
		ID:   c.AggregateID,
		Name: command.Name,
	}

	// 檢查User
	if err := c.userService.CheckUserData(ctx, user); err != nil {
		return nil, err
	}

	// 建立User
	if err := c.userRepo.CreateUser(ctx, user); err != nil {
		return nil, err
	}

	//fmt.Println("AggregateID", c.AggregateID)
	return &UserCreated{ID: c.AggregateID}, nil
}
