package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
	"time"
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
	user := NewUser(conn, s)
	user.Online()

	// 监听用户是否活跃的channel
	isLive := make(chan bool)

	//接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf) // 从conn里面读
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn read err:", err)
				return
			}
			msg := string(buf[:n-1]) //去除最后的\n

			user.DoMessage(msg)

			// 用户活跃
			isLive <- true
		}
	}()

	// 当前Handle阻塞
	for {
		select {
		case <-isLive:
			// 当前用户活跃，应该重置定时器
			// 不做任何事情，为了激活select，更新定时器
		case <-time.After(time.Second * 20): // 一旦执行这个case，会重置定时器
			user.SendMsg("强制下线！\n")
			// user.Offline() 为什么这里不需要Offline? 答:上面的for循环检测到n==0，会自动调用Offline

			// 销毁资源
			close(user.C)
			conn.Close()
			runtime.Goexit()
		}
	}

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
