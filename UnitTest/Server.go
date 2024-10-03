// zinx/UnitTest/Server.go
package main

import (
	"zinx/znet"
)

// Server模块的测试函数
func main() {

	// 1.创建一个Server句柄s
	s := znet.NewServer("[zinx v0.1]")

	// 2.开启服务
	s.Serve()
}
