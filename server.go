package main

import (
	"github.com/tidwall/redcon"
	log "github.com/sirupsen/logrus"
	"github.com/liad344/go-redis/redis"
)

type handleConnection func(conn redcon.Conn) bool
type handleClosedConnection func(conn redcon.Conn, err error)

type ServerConfig struct {
	addr string
}

type Server struct {
	cfg    ServerConfig
	ins    *redis.Instance
	mux    *redcon.ServeMux
	accept handleConnection
	closed handleClosedConnection
}


func (s *Server) Start(){
	s.Init()
	s.ServeHTTP()
}
func (s *Server) ServeHTTP() {
	log.Info("Serving redis clone")

	if err := redcon.ListenAndServe(":8000" , s.mux.ServeRESP , s.accept , s.closed); err != nil {
		log.Error("Could not start http server")
	}
}
func (s *Server) Init(){
	s.mux = redcon.NewServeMux()
	s.ins = redis.NewInstance()
	s.accept = onNewConnection
	s.closed = onConnectionClosed
	s.mux.HandleFunc("set" , s.ins.Set)
	s.mux.HandleFunc("get" , s.ins.Get)
	s.mux.HandleFunc("del" , s.ins.Del)
	s.mux.HandleFunc("ping" , s.ins.Ping)
}

func NewServer() *Server{
	return &Server{
		cfg:    ServerConfig{},
		ins:    nil,
		mux:    nil,
		accept: nil,
		closed: nil,
	}
}