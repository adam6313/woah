package user

import "woah/internal/service/user/usecase/create"

// CreateUser -
type CreateUser struct {
	Name string `json:"name"`
}

// AggregateID -
func (c *CreateUser) AggregateID() string {
	return ""
}

// Message -
func (c *CreateUser) Message() interface{} {
	return &create.CreateUser{
		Name: c.Name,
	}
}
