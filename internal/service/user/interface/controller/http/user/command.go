package user

import (
	"woah/internal/service/user/usecase/create"

	"woah/internal/service/user/usecase/update"
)

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

// UpdateUserInfo -
type UpdateUserInfo struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// AggregateID -
func (u *UpdateUserInfo) AggregateID() string {
	return u.ID
}

// Message -
func (u *UpdateUserInfo) Message() interface{} {
	return &update.UpdateUserInfo{
		Name: u.Name,
	}
}
