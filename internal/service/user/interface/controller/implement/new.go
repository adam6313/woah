package implement

import (
	"woah/internal/common/command"
	"woah/internal/service/user/usecase/create"
	"woah/internal/service/user/usecase/update"
)

// New -
func New() command.Dispatch {
	return command.NewDispatch(
		create.NewUseCase(),
		update.NewUseCase(),
	)
}
