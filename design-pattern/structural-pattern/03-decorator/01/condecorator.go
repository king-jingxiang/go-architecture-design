package _1

import "fmt"

// DecoratorA 具体装饰器A
type DecoratorA struct {
	Decorator
}

// Operation 具体装饰器A的实现
func (d *DecoratorA) Operation() string {
	operation := d.component.Operation()
	return d.IndependentMethod(operation)
}

// IndependentMethod 具体装饰器A的独有方法
func (d *DecoratorA) IndependentMethod(operation string) string {
	return fmt.Sprintf("%s-%s", operation, "DecoratorA")
}

// DecoratorB 具体装饰器A
type DecoratorB struct {
	Decorator
}

// Operation 具体装饰器A的实现
func (d *DecoratorB) Operation() string {
	operation := d.component.Operation()
	return d.IndependentMethod(operation)
}

// IndependentMethod 具体装饰器A的独有方法
func (d *DecoratorB) IndependentMethod(operation string) string {
	return fmt.Sprintf("%s-%s", operation, "DecoratorB")
}
