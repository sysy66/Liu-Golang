// Client.go
package main

import (
	"fmt"
	"net"
	"zinx/znet"
)

func main() {

	// 客户端 goroutine ，负责模拟黏包的数据，然后进行发送
	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	// 1.创建一个封包对象 dp
	dp := znet.NewDataPack()

	// 2.封装一个 msg1 包
	msg1 := &znet.Message{
		Id:      0,
		DataLen: 5,
		Data:    []byte{'h', 'e', 'l', 'l', 'o'},
	}
	sendData1, err := dp.Pack(msg1)
	if err != nil {
		fmt.Println("Client pack data err:", err)
		return
	}

	// 3.封装一个 msg2 包
	msg2 := &znet.Message{
		Id:      1,
		DataLen: 7,
		Data:    []byte{'w', 'o', 'r', 'l', 'd', '!', '!'},
	}
	sendData2, err := dp.Pack(msg2)
	if err != nil {
		fmt.Println("Client pack data err:", err)
		return
	}

	// 4.将 sendData1 和 sendData2 拼接到一起，组成黏包
	sendData1 = append(sendData1, sendData2...)

	// 5.向服务器端写数据
	conn.Write(sendData1)

	// 客户端阻塞
	select {}
}
