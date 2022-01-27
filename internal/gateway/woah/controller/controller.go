package controller

import (
	"woah/internal/common/command"
	"woah/internal/gateway/woah/controller/user"
	"woah/internal/service/user/interface/controller/implement"

	"github.com/kataras/iris/v12"
)

func New() {
	app := iris.New()

	dispatch := mergeDispatch()

	setController(app, dispatch)

	app.Run(iris.Addr(":3000"))
}

func mergeDispatch() command.Dispatch {
	return command.Merger(
		implement.New(),
	)
}

func setController(app *iris.Application, dispatch command.Dispatch) {
	user.NewController(dispatch).SetRouter(app)
}
