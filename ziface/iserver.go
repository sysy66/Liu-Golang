// Package ziface zinx/ziface/iserver.go
package ziface

// IServer 定义服务器接口
type IServer interface {
	// Start 启动服务器方法
	Start()
	// Stop 通知服务器方法
	Stop()
	// Serve 开启业务服务方法
	Serve()
}
