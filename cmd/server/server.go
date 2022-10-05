package server

import (
	"github.com/albertopformoso/inventory/internal/routes"
	"github.com/labstack/echo/v4"
)

type Server struct {
	routes routes.Routes
}

func New(routes routes.Routes) *Server {
	return &Server{
		routes: routes,
	}
}

func (s *Server) Start(e *echo.Echo, address string) error {
	s.routes.Init(e)
	return e.Start(address)
}