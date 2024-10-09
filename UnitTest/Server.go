// zinx/UnitTest/Server.go
package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

// PingRouter ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter // 一定要先定义基础路由 BaseRouter
}

// Handle Test
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Server模块的测试函数
func main() {

	// 1.创建一个Server句柄s
	s := znet.NewServer()

	// 2.配置路由
	s.AddRouter(&PingRouter{})

	// 3.开启服务
	s.Serve()
}
