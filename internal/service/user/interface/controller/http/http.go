package http

import (
	"net/http"
	"woah/internal/common/command"
	"woah/internal/service/user/interface/controller/http/user"
	"woah/internal/service/user/usecase/create"
	"woah/internal/service/user/usecase/update"

	"github.com/kataras/iris/v12"
)

// NewHTTPServer -
func NewHTTPServer(createUserUsecase create.CreateUserUsecase, updateUserUsecase update.UpdateUserUsecase) http.Handler {
	h := user.Server{
		App: iris.New(),
		Dispatch: command.NewDispatch(
			createUserUsecase,
			updateUserUsecase,
		),
	}

	h.SetRouter()

	return h.App
}
