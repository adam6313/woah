package main

import (
	"woah/internal/service/user/cmd"
)

type User struct {
	Name string
}

func main() {
	cmd.Execute()

	//u := &User{
	//Name: "adam",
	//}

	//c := NewCommand("123", u)

	//fmt.Println(c.Command())

	//handle(c)
}

// Command -
//type Command interface {
//// AggregateID -
//AggregateID() string

//// CommandType -
//Type() string

//// Command -
//Command() interface{}
//}

//// CommandDescriptor -
//type CommandDescriptor struct {
//id string

//command interface{}
//}

//// NewCommand -
//func NewCommand(aggregateID string, command interface{}) Command {
//return &CommandDescriptor{
//id:      aggregateID,
//command: command,
//}
//}

//// CommandType -
//func (c *CommandDescriptor) Type() string {
//return reflect.TypeOf(c).Elem().Name()
//}

//// AggregateID -
//func (c *CommandDescriptor) AggregateID() string {
//return c.id
//}

//// Command -
//func (c *CommandDescriptor) Command() interface{} {
//return c.command
//}

//func handle(message Command) {
//switch cmd := message.Command().(type) {
//case *User:
//fmt.Println(cmd.Name)

//fmt.Println(message.AggregateID())
//}
//}
