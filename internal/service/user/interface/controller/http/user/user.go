package user

import (
	"fmt"
	"woah/internal/common/command"
	"woah/internal/service/user/interface/controller/http/middle"
	"woah/internal/service/user/usecase/create"
	"woah/internal/service/user/usecase/login"
	"woah/internal/service/user/usecase/update"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

// server -
type Server struct {
	App               *iris.Application
	CreateUserUsecase create.CreateUserUsecase
	UpdateUserUsecase update.UpdateUserUsecase
	Dispatch          command.Dispatch
}

// CreateUser -
func (s *Server) CreateUser(c *middle.C) {
	aggregateID := uuid.New().String()

	// new command
	cmd := command.New(aggregateID, &create.CreateUser{
		Name: "adam",
	})

	// dispatch command handle
	result, err := s.Dispatch.Handle(c.Request().Context(), cmd)
	if err != nil {
		c.E(err)
		return
	}

	r := result.(*create.UserCreated)

	fmt.Println(r)

	c.R(aggregateID)
}

// UpdateUser -
func (s *Server) UpdateUser(c *middle.C) {
	aggregateID := c.Params().Get("id")

	data := new(update.UpdateUserInfo)
	if err := c.ReadJSON(&data); err != nil {
		return
	}

	// new command
	cmd := command.New(aggregateID, &update.UpdateUserInfo{
		Name: data.Name,
	})

	// dispatch command handle
	_, err := s.Dispatch.Handle(c.Request().Context(), cmd)
	if err != nil {
		c.E(err)
		return
	}

	//r := result.(*create.UserCreated)

	c.R(aggregateID)
}

// LoginUser -
func (s *Server) Login(c *middle.C) {
	//body, _ := c.GetBody()

	aggregateID := c.Params().Get("id")

	data := new(login.UserLoginCmd)

	if err := c.ReadJSON(data); err != nil {
		c.E(err)
		return
	}

	// new command
	cmd := command.New(aggregateID, data)

	// dispatch command handle
	result, err := s.Dispatch.Handle(c.Request().Context(), cmd)
	if err != nil {
		c.E(err)
		return
	}

	//r := result.(*create.UserCreated)

	c.R(result.(*login.UserLogined))
}
