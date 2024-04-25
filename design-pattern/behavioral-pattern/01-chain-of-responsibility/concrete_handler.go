package _1_chain_of_responsibility

import "fmt"

type ConcreteHandler struct {
}

func (c *ConcreteHandler) Handle(handlerId int) string {
	return fmt.Sprintf("ConcreteHandler %d", handlerId)
}
