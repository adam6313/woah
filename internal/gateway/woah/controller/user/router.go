package user

import (
	"woah/internal/gateway/woah/middle"

	"github.com/kataras/iris/v12"
)

// SetRouter -
func (ct *controller) SetRouter(app *iris.Application) {

	// Create user
	app.Post("/api/user", middle.HandleFunc(ct.command(new(CreateUser))))
}
