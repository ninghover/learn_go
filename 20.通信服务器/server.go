package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int
	// 在线用户列表
	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	//消息广播的channel
	Message chan string
}

// 创建一个Server
func NewServer(ip string, port int) *Server {
	server := Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return &server
}

// 监听Message广播消息channel，一旦有消息就发送给所有的User (单独启一个goroutine)
func (s *Server) ListenMessage() {
	for {
		msg := <-s.Message
		s.mapLock.Lock()
		for _, cli := range s.OnlineMap {
			cli.C <- msg
		}
		s.mapLock.Unlock()
	}
}

// 广播消息
func (s *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + msg
	s.Message <- sendMsg
}

func (s *Server) Handle(conn net.Conn) {
	// 用户上线，将用户加入到在线用户表中
	user := NewUser(conn)
	s.mapLock.Lock()
	s.OnlineMap[user.Name] = user
	s.mapLock.Unlock()

	// 广播用户上线消息
	s.BroadCast(user, "已上线")

	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				s.BroadCast(user, "下线了")
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn read err:", err)
				return
			}
			msg := string(buf[:n-1]) //去除最后的\n
			
			s.BroadCast(user, msg)
		}
	}()

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

	// 启动监听Message的goroutine
	go s.ListenMessage()

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
