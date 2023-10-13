package utils

import (
	"TCP_Framework/zinx/ziface"
	"encoding/json"
	"os"
)

type GlobalObj struct {
	TcpServer     ziface.IServer
	Host          string
	TcpPort       int
	Name          string
	Version       string
	MaxpacketSize uint32
	MaxConn       int
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {

	data, err := os.ReadFile("zinx/conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:          "ZinxServerApp",
		Version:       "V0.4",
		TcpPort:       7777,
		Host:          "0.0.0.0",
		MaxConn:       12000,
		MaxpacketSize: 4096,
	}
	GlobalObject.Reload()
}
