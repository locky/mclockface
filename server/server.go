package server

import "fmt"

type Server struct{}

func (srv *Server) Lock() error {
	return fmt.Errorf("not implemented")
}

func (srv *Server) Unlock() error {
	return fmt.Errorf("not implemented")
}
