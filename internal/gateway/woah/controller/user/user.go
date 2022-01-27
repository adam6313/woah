package user

import (
	"woah/internal/common/command"
	"woah/internal/gateway/woah/middle"
)

type controller struct {
	Dispatch command.Dispatch
}

// NewController -
func NewController(dispatch command.Dispatch) *controller {
	return &controller{
		Dispatch: dispatch,
	}
}

// Command -
func (ct *controller) command(cmd command.Command) func(c *middle.C) {
	return func(c *middle.C) {
		if err := c.ReadJSON(cmd); err != nil {
			c.E(err)
		}

		// new command
		cmd := command.New(cmd.AggregateID(), cmd.Message())

		// dispatch command handle
		result, err := ct.Dispatch.Handle(c.Request().Context(), cmd)
		if err != nil {
			c.E(err)
			return
		}

		c.R(result)
	}
}
