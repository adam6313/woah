package user

import (
	"woah/internal/service/user/interface/controller/http/middle"
)

func (s *Server) SetRouter() {
	r := s.App.Party("/api/user")

	{
		// 建立會員
		r.Post("", middle.HandleFunc(s.CreateUser))

		// 更新會員
		r.Put("", middle.HandleFunc(s.UpdateUser))
	}
}
