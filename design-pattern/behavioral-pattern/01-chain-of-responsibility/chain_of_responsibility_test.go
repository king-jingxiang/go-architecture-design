package _1_chain_of_responsibility

import (
	"fmt"
	"testing"
)

func TestBaseHandler_Handle(t *testing.T) {
	handler1 := NewBaseHandler("Test01", nil, 1)
	handler2 := NewBaseHandler("Test01", handler1, 2)
	handler3 := NewBaseHandler("Test01", handler2, 3)

	handleId1 := handler1.Handle(1)
	handler1.SetNext(handler2)

	handleId2 := handler1.Handle(handleId1)
	fmt.Println(handleId1, handleId2)

	handler2.SetNext(handler3)
	handleId3 := handler2.Handle(handleId2)

	fmt.Println(handleId3)
}
