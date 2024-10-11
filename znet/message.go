package znet

type Message struct {
	Id      uint32 // 消息的ID
	DataLen uint32 // 消息的长度
	Data    []byte // 消息的内容
}

// NewMessage 创建一个 Message 消息包
func NewMessage(id uint32, data []byte) *Message {
	return &Message{
		Id:      id,
		DataLen: uint32(len(data)),
		Data:    data,
	}
}

// GetMsgId 获取消息的ID
func (msg *Message) GetMsgId() uint32 {
	return msg.Id
}

// GetDataLen 获取消息的数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

// GetData 获取消息的内容
func (msg *Message) GetData() []byte {
	return msg.Data
}

// SetMsgId 设计消息ID
func (msg *Message) SetMsgId(msgId uint32) {
	msg.Id = msgId
}

// SetDataLen 设置消息数据段长度
func (msg *Message) SetDataLen(len uint32) {
	msg.DataLen = len
}

// SetData 设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.Data = data
}
