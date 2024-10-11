// zinx/UnitTest/Client.go
package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/znet"
)

func main() {

	fmt.Println("Client Test ... start")
	// 3s之后发起测试请求，给服务器端开启服务的机会
	time.Sleep(3 * time.Second)

	conn, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		// 发送封包 message 消息
		dp := znet.NewDataPack()
		msg, _ := dp.Pack(znet.NewMessage(0, []byte("Zinx V0.5 Client Test Message")))
		_, err := conn.Write(msg)
		if err != nil {
			fmt.Println("client Write err, exit!", err)
			return
		}

		// 先读出流中的 head 部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) // ReadFull 会把 msg填充满
		if err != nil {
			fmt.Println("client Read err: ", err)
			break
		}

		//  将 headData 字节流拆包到 msg 中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("client Unpack err: ", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			// msg 有 data 数据，需要再次读取 data 数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			// 根据 dataLen 从 io 中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("client Read err: ", err)
				return
			}

			fmt.Println("==> Recv Msg: ID = ", msg.GetMsgId(), ", len = ", msg.GetDataLen(), ", data = ", string(msg.GetData()))
		}

		time.Sleep(1 * time.Second)
	}
}
