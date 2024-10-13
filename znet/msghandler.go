package znet

import (
	"fmt"
	"zinx/ziface"
)

type MsgHandle struct {
	Apis map[uint32]ziface.IRouter // 存放每个 MsgId 所对应的处理方法的 map 属性
}

func NewMsgHandle() *MsgHandle {
	return &MsgHandle{
		Apis: make(map[uint32]ziface.IRouter),
	}
}

// DoMsgHandler 以非阻塞方式处理消息
func (mh *MsgHandle) DoMsgHandler(request ziface.IRequest) {
	handler, ok := mh.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msgId = ", request.GetMsgId(), "is not FOUND!")
		return
	}
	// 执行对应的处理方法
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}

// AddRouter 为消息添加具体地处理逻辑
func (mh *MsgHandle) AddRouter(msgId uint32, router ziface.IRouter) {
	// 1.判断当前 msg 绑定的 API 处理方法是否已经存在
	if _, ok := mh.Apis[msgId]; ok {
		panic("repeated api , msgId = " + fmt.Sprint(msgId))
	}
	// 2.添加 msg 与 api 的绑定关系
	mh.Apis[msgId] = router
	fmt.Println("Add api msgId = ", msgId)
}
