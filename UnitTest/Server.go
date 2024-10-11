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
	fmt.Println("Call PingRouter Handle")
	// 先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client: msgId = ", request.GetMsgId(), ", data = ", string(request.GetData()))

	// 回写数据
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping\n"))
	if err != nil {
		fmt.Println(err)
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
