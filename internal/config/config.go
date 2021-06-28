package config

import "fmt"

// Server stuct describe params of server
type Server struct {
	Port string `env:"PORT" env-default:"8081"`
	Host string `env:"HOST" env-default:"0.0.0.0"`
}

func (s *Server) String() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
