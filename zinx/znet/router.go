package znet

import "TCP_Framework/zinx/ziface"

type BaseRouter struct{}

func (br *BaseRouter) PreHandle(res ziface.IRequest)  {}
func (br *BaseRouter) Handle(req ziface.IRequest)     {}
func (br *BaseRouter) PostHandle(req ziface.IRequest) {}
