package websocket

import (
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
)

type Server struct {
	addr     string
	upgrader websocket.Upgrader
	logx.Logger
}

func NewServer(addr string) *Server {
	return &Server{
		addr:     addr,
		upgrader: websocket.Upgrader{},
		Logger:   logx.WithContext(context.Background()),
	}
}

func (s *Server) ServerWs(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			s.Errorf("server handler ws recover err %v", r)
		}
	}()
}

func (s *Server) Start() {
	http.HandleFunc("/ws", s.ServerWs)
	fmt.Println("websocket server start")
	s.Info(http.ListenAndServe(s.addr, nil))
}

func (s *Server) Stop() {
	fmt.Println("websocket server stop")
}
