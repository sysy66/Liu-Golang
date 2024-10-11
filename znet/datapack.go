package znet

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/utils"
	"zinx/ziface"
)

// DataPack 封包拆包类实例，暂不需要成员
type DataPack struct{}

// NewDataPack 封包拆包实例初始化方法
func NewDataPack() *DataPack {

	return &DataPack{}
}

// GetHeadLen 获取包头长度办法
func (dp *DataPack) GetHeadLen() uint32 {
	// Id uint32（4字节）+ DataLen uint32（4字节）
	return 8
}

// Pack 封包方法
func (dp *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {

	// 创建一个存放 Bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	// 写 dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return nil, err
	}

	// 写 msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetMsgId()); err != nil {
		return nil, err
	}

	// 写 data
	if err := binary.Write(dataBuff, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}

	return dataBuff.Bytes(), nil
}

// Unpack 拆包方法（解压数据）
func (dp *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {

	// 创建一个输入二进制数据的 ioReader
	dataBuff := bytes.NewReader(binaryData)

	//只解压 head 的信息，得到 dataLen 和 msgID
	msg := &Message{}

	// 读 dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.DataLen); err != nil {
		return nil, err
	}

	// 读 msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &msg.Id); err != nil {
		return nil, err
	}

	// 判断 dataLen 是否超过允许的最大包长度
	if utils.GlobalObject.MaxPacketSize > 0 && msg.DataLen > utils.GlobalObject.MaxPacketSize {
		return nil, errors.New("too large msg data received")
	}

	//这里只需要把 head 的数据拆包出来，然后通过 head 的长度，再从 conn 读取一次数据
	return msg, nil
}
