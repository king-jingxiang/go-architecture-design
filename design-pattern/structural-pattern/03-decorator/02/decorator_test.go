package _1

import (
	"testing"
)

func TestDecorator_Operation2(t *testing.T) {
	concreteComponent := &ConComponentV1{}
	decoratorA := &DecoratorWrapper{HandlerFunc: func(s string) string {
		return s + "-DecoratorA"
	}}
	decoratorB := &DecoratorWrapper{HandlerFunc: func(s string) string {
		return s + "-DecoratorB"
	}}
	decoratorA.SetComponent(concreteComponent)
	r1 := decoratorA.Operation()
	if r1 != "ConComponent Operation V1-DecoratorA" {
		t.Errorf("Expected result to be 'ConComponent Operation V1-DecoratorA', got '%s'", r1)
	}
	decoratorB.SetComponent(decoratorA)
	r2 := decoratorB.Operation()
	if r2 != "ConComponent Operation V1-DecoratorA-DecoratorB" {
		t.Errorf("Expected result to be 'ConComponent Operation V1-DecoratorA-DecoratorB', got '%s'", r2)
	}
}
