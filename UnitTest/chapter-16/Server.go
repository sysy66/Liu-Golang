// Server.go
package main

import (
	"fmt"
	"io"
	"net"
	"zinx/znet"
)

// 只是负责测试 datapack 拆包和封包功能
func main() {

	// 创建 socket TCP Server
	listener, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("server listen err:", err)
		return
	}

	// 创建服务器 goroutine ，负责从客户端 goroutine 读取黏包的数据，然后进行解析
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server accept err:", err)
		}

		// 处理客户端请求
		go func(conn net.Conn) {
			// 创建封包拆包对象 dp
			dp := znet.NewDataPack()

			for {
				// 1.先读取流中的 head 部分
				headData := make([]byte, dp.GetHeadLen())
				// ReadFull 会把 msg 填满为止
				_, err := io.ReadFull(conn, headData)
				if err != nil {
					fmt.Println("read head err")
					break
				}

				// 2.将 headData 字节流拆包到 msg 中
				msgHead, err := dp.Unpack(headData)
				if err != nil {
					fmt.Println("server unpack err:", err)
					return
				}

				// 3.根据 dataLen 从 io 中读取字节流
				if msgHead.GetDataLen() > 0 {
					// msg 有 data 数据，需要再次读取 data 数据
					msg := msgHead.(*znet.Message)
					msg.Data = make([]byte, msg.GetDataLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						fmt.Println("server unpack data err:", err)
						return
					}

					fmt.Println("==> Recv Msg: ID=", msg.Id, ", Len=", msg.DataLen, ", data=", string(msg.Data))
				}
			}
		}(conn)
	}
}
