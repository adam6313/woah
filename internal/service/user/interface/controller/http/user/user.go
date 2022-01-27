package user

import (
	"woah/internal/common/command"
	"woah/internal/service/user/interface/controller/http/middle"

	"github.com/kataras/iris/v12"
)

// server -
type Server struct {
	App      *iris.Application
	Dispatch command.Dispatch
}

// Command -
func (s *Server) command(cmd command.Command) func(c *middle.C) {
	return func(c *middle.C) {
		if err := c.ReadJSON(cmd); err != nil {
			c.E(err)
		}

		// new command
		cmd := command.New(cmd.AggregateID(), cmd.Message())

		// dispatch command handle
		result, err := s.Dispatch.Handle(c.Request().Context(), cmd)
		if err != nil {
			c.E(err)
			return
		}

		c.R(result)
	}
}
