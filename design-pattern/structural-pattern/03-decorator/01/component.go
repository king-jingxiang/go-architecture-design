package _1

// Component 组件接口
type Component interface {
	Operation() string
}

// ConComponentV1 具体组件
type ConComponentV1 struct {
}

// Operation 具体组件方法
func (c *ConComponentV1) Operation() string {
	return "ConComponent Operation V1"
}

// ConComponentV2 具体组件
type ConComponentV2 struct {
}

// Operation 具体组件方法
func (c *ConComponentV2) Operation() string {
	return "ConComponent Operation V2"
}
