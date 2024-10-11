// Package znet zinx/znet/connection.go
package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"zinx/ziface"
)

type Connection struct {
	// 当前连接的stocket TCP套接字
	Conn *net.TCPConn
	// 当前连接的ID，也可以称为SessionID，ID全局唯一
	ConnID uint32
	// 当前连接的关闭状态
	isClosed bool
	// 该连接的处理方法 router
	Router ziface.IRouter
	// 告知该连接已经退出/停止的channel
	ExitBuffChan chan bool
}

// NewConnection 创建连接的方法
func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:         conn,
		ConnID:       connID,
		isClosed:     false,
		Router:       router,
		ExitBuffChan: make(chan bool, 1),
	}

	return c
}

// StartReader 处理conn读数据的Goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), "conn reader exit!")
	defer c.Stop()

	for {
		// 创建封包拆包对象 dp
		dp := NewDataPack()

		// 读取客户端的 Msg head
		headData := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(c.GetTCPConnection(), headData); err != nil {
			fmt.Println("read msg head data error:", err)
			c.ExitBuffChan <- true
			continue
		}

		// 拆包，得到 msgId 和 dataLen 后放到 msg 中
		msg, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("unpack msg error:", err)
			c.ExitBuffChan <- true
			continue
		}

		// 根据 dataLen 读取 data，放到 msg.Data 中
		var data []byte
		if msg.GetDataLen() > 0 {
			data = make([]byte, msg.GetDataLen())
			if _, err := io.ReadFull(c.GetTCPConnection(), data); err != nil {
				fmt.Println("read msg data error:", err)
				c.ExitBuffChan <- true
				continue
			}
			msg.SetData(data)
		}

		// 得到当前客户端请求的 Request 数据
		req := Request{
			conn: c,
			msg:  msg, // 将之前的 buf 改成 msg
		}
		// 从路由 Routers中找到注册绑定 Conn 的对应 Handle
		go func(request ziface.IRequest) {
			// 执行注册的路由方法
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}

// Start 启动连接，让当前连接开始工作
func (c *Connection) Start() {

	// 开启处理该连接读取客户端数据之后的请求业务
	go c.StartReader()

	for {
		select {
		case <-c.ExitBuffChan:
			// 得到退出消息，不再阻塞
			return
		}
	}
}

// Stop 停止连接，结束当前连接状态
func (c *Connection) Stop() {
	// 1.如果当前连接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	// TODO Connection Stop() 如果用户注册了该连接的关闭回调业务，则在此刻应该显示调用

	// 关闭 Stocket 连接
	c.Conn.Close()

	// 通知从缓冲队列读取数据的业务，该连接已经关闭
	c.ExitBuffChan <- true

	// 关闭该连接的全部管道
	close(c.ExitBuffChan)
}

// GetTCPConnection 从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// RemoteAddr 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// SendMsg  直接将 Message 数据发送给远程的 TCP 客户端
func (c *Connection) SendMsg(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection is closed when send msg")
	}

	// 将 data 封包，并且发送
	dp := NewDataPack()
	msg, err := dp.Pack(NewMessage(msgId, data))
	if err != nil {
		fmt.Println("pack error msg id:", msgId)
		return errors.New("pack error msg")
	}

	// 写回客户端
	if _, err := c.Conn.Write(msg); err != nil {
		fmt.Println("write error msg id:", msgId)
		c.ExitBuffChan <- true
		return errors.New("conn Write error")
	}

	return nil
}
