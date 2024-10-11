package ziface

// IMessage 将请求的一条消息封装到 message 中，定义抽象层接口
type IMessage interface {
	GetMsgId() uint32   // 获取消息ID
	GetDataLen() uint32 // 获取消息数据段长度
	GetData() []byte    // 获取消息内容

	SetMsgId(uint32)   // 设计消息ID
	SetDataLen(uint32) // 设置消息数据段长度
	SetData([]byte)    // 设计消息内容
}
