package po

import (
	"woah/internal/service/user/domain/model/aggregate"
)

func ConverToUserPo(in *aggregate.User) *UserPo {

	return &UserPo{
		ID:   in.ID,
		Name: in.Name,
	}

}
