package _1_chain_of_responsibility

import "fmt"

// Handler 接口定义
type Handler interface {
	SetNext(handler Handler)
	Handle(handlerId int) int
}

type BaseHandler struct {
	name      string
	next      Handler
	handlerId int
}

func NewBaseHandler(name string, next Handler, handlerId int) Handler {
	return &BaseHandler{
		name:      name,
		next:      next,
		handlerId: handlerId,
	}
}

func (h *BaseHandler) SetNext(handler Handler) {
	h.next = handler
}

func (h *BaseHandler) Handle(handlerId int) int {
	if handlerId < 4 {
		ch := &ConcreteHandler{}
		resp := ch.Handle(handlerId)
		fmt.Println(resp)
		if h.next != nil {
			h.next.Handle(handlerId + 1)
		}
		return handlerId + 1
	}
	return 0
}
