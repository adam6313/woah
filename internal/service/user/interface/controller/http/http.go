package http

import (
	"context"
	"net/http"
	"woah/internal/common/command"
	"woah/internal/service/user/usecase/create"

	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type HTTPServer struct {
	App      *iris.Application
	u        create.CreateUserUsecase
	dispatch command.Dispatch
}

// NewHttpServer -
func NewHttpServer(u create.CreateUserUsecase) http.Handler {
	h := HTTPServer{
		App:      iris.New(),
		u:        u,
		dispatch: command.NewDispatch(u),
	}

	h.App.Get("/user", h.createTest)

	return h.App
}

func (h HTTPServer) createTest(c iris.Context) {

	aggregateID := uuid.New().String()

	cmd := command.NewCommand(aggregateID, &create.CreateUser{
		Name: "adam",
	})

	h.dispatch.Handle(context.Background(), cmd)
}
