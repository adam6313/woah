package command

import "reflect"

// Command -
type Command interface {
	// AggregateID -
	AggregateID() string

	// CommandType -
	Type() (reflect.Type, string)

	// Command -
	Command() interface{}
}

// CommandDescriptor -
type CommandDescriptor struct {
	id string

	command interface{}
}

// NewCommand -
func NewCommand(aggregateID string, command interface{}) Command {
	return &CommandDescriptor{
		id:      aggregateID,
		command: command,
	}
}

// CommandType -
func (c *CommandDescriptor) Type() (reflect.Type, string) {
	t := reflect.TypeOf(c.command)

	return t, t.Name()
}

// AggregateID -
func (c *CommandDescriptor) AggregateID() string {
	return c.id
}

// Command -
func (c *CommandDescriptor) Command() interface{} {
	return c.command
}
