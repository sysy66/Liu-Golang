package ziface

// IDataPack 封包数据和拆包数据
// 直接面向 TCP 连接中的数据流，为传输数据添加头部信息，用于处理 TCP 黏包问题
type IDataPack interface {
	GetHeadLen() uint32               			// 获取包头长度办法
	Pack(msg IMessage) ([]byte, error) 			// 封包方法
	Unpack(msg []byte) (IMessage, error) 		// 拆包方法
}