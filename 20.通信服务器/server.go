package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个Server

func NewServer(ip string, port int) *Server {
	server := Server{
		Ip:   ip,
		Port: port,
	}
	return &server
}

func (s *Server) Handle(conn net.Conn) {
	fmt.Println("连接建立成功...")
}

// 启动服务器的接口
func (s *Server) Start() {
	//socket listen
	Listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Println("net.Listen err: ", err)
		return
	}

	//close
	defer Listener.Close()

	for {
		//accept
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}

		//do handle
		go s.Handle(conn)
	}
}
