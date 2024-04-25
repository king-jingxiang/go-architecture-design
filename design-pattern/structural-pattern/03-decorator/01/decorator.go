package _1

// Decorator 装饰器
type Decorator struct {
	component Component
}

// SetComponent 设置组件
func (d *Decorator) SetComponent(c Component) {
	d.component = c
}

// Operation 装饰方法
func (d *Decorator) Operation() string {
	if d.component != nil {
		return d.component.Operation()
	}
	return ""
}
