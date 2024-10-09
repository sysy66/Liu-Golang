// Package znet zinx/znet/server.go
package znet

import (
	"fmt"
	"net"
	"zinx/utils"
	"zinx/ziface"
)

// Server IServer接口实现，定义一个Server服务类
type Server struct {
	// 服务器的名称
	Name string
	// tcp4 or other
	IPVersion string
	// 服务绑定的IP地址
	IP string
	// 服务绑定的端口
	Port int
	// 当前 Server 由用户绑定回调 router，也就是 Server 注册的连接对应的处理业务
	Router ziface.IRouter
}

// ====================== 实现 ziface.IServer 里的全部接口方法 ======================

// Start 开启网络服务
func (s *Server) Start() {
	fmt.Printf("[START] Server name: %s, listener at IP: %s, Port %d, is starting\n", s.Name, s.IP, s.Port)
	fmt.Printf("[Zinx] Version: %s, MaxConn: %d, MaxPacketSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPacketSize)

	// 开启一个go去做服务器端Listener业务
	go func() {
		// 1.获取一个TCP的Addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			return
		}

		// 2.监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen", s.IPVersion, "err", err)
			return
		}

		// 已经监听成功
		fmt.Println("start Zinx server", s.Name, "succ, now listening...")

		// TODO server.go 应该有一个自动生成ID的方法
		var cid uint32
		cid = 0

		// 3.启动server网络连接业务
		for {
			// 3.1阻塞等待客户端建立连接请求
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err", err)
				continue
			}

			// 3.2 TODO Server.Start() 设置服务器最大连接控制，如果超过最大连接，则关闭新的连接
			// 3.3 TODO Server.Start() 处理该新连接请求的业务方法，此时handler和conn应该是绑定的
			dealConn := NewConnection(conn, cid, s.Router)
			cid++

			// 3.4 启动当前连接的处理业务
			go dealConn.Start()

		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx Server, name", s.Name)

	// TODO Server.Stop() 将需要清理的连接信息或者其他信息一并停止或者清理
}

func (s *Server) Serve() {
	s.Start()

	// TODO Server.Serve() 如果在启动服务的时候还要处理其他的事情，则可以在这里添加

	// 阻塞，否则主Go，listener的go将会退出
	select {}
}

// AddRouter 路由功能：给当前服务注册一个路由业务方法，供客户端连接处理使用
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router

	fmt.Println("Add Router succ!")
}

// NewServer 创建一个服务器句柄
func NewServer() ziface.IServer {
	// 先初始化全局配置文件
	utils.GlobalObject.Reload()
	s := &Server{
		Name:      utils.GlobalObject.Name, // 从全局参数获取
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,    // 从全局参数获取
		Port:      utils.GlobalObject.TcpPort, // 从全局参数获取
		Router:    nil,                        // 默认不制指定
	}
	return s
}
