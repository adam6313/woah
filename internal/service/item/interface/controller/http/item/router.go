package item

import "woah/internal/service/user/interface/controller/http/middle"

// SetRouter -
func (s *Server) SetRouter() {
	r := s.App.Party("/api/item")
	{
		// 建立會員
		r.Post("", middle.HandleFunc(s.Warehouse))

		// 上架/下架
		r.Put("/:id/:status", middle.HandleFunc(s.Shelf))
	}
}
