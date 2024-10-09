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

// PreHandle Test
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping...\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Handle Test
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// PostHandle Test
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")

	_, err := request.GetConnection().GetTCPConnection().Write([]byte("after ping...\n"))
	if err != nil {
		fmt.Println("call back ping ping ping error")
	}
}

// Server模块的测试函数
func main() {

	// 1.创建一个Server句柄s
	s := znet.NewServer("[zinx v0.3]")

	// 2.开启服务
	s.AddRouter(&PingRouter{})

	// 3.开启服务
	s.Serve()
}
