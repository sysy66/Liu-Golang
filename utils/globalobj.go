package utils

import (
	"encoding/json"
	"io/ioutil"
	"zinx/ziface"
)

// GlobalObj 存储一切有关 Zinx 框架的全局参数，供其他模块使用
// 一些参数也可以通过用户根据 zinx.json 来配置
type GlobalObj struct {
	TcpServer ziface.IServer // 当前 Zinx 的全局 Server 对象
	Host      string         // 当前服务器主机 IP
	TcpPort   int            // 当前服务器主机监听端口号
	Name      string         // 当前服务器名称
	Version   string         // 当前 Zinx 版本号

	MaxPacketSize uint32 // 读取数据包的最大值
	MaxConn       int    // 当前服务器主机允许的最大连接个数
}

// GlobalObject 定义一个全局的对象
var GlobalObject *GlobalObj

// Reload 读取用户的配置文件
func (g *GlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	// 将 json 数据解析到 struct 中
	// fmt.Printf("json:%s\n", data)
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

// 提供 init 方法，默认加载
func init() {
	// 初始化 GlobalObject 变量，设置一些默认值
	GlobalObject = &GlobalObj{
		Name:          "ZinxServerApp",
		Version:       "V0.4",
		TcpPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxPacketSize: 4096,
	}

	// 从配置文件中加载一些用户配置的参数
	GlobalObject.Reload()
}
