package _1

type HandlerFunc func(string) string

// DecoratorWrapper 具体装饰器
type DecoratorWrapper struct {
	Decorator
	HandlerFunc HandlerFunc
}

// Operation 具体装饰器A的实现
func (d *DecoratorWrapper) Operation() string {
	return d.HandlerFunc(d.component.Operation())
}
