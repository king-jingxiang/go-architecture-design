package _3_abstract_factory

// AbstractProduct 抽象产品接口
type AbstractProduct interface {
	GetName() string
}

// ConcreteProductV1 具体产品
type ConcreteProductV1 struct {
}

// GetName 具体产品的实现方法
func (c *ConcreteProductV1) GetName() string {
	return "ConcreteProductV1"
}

// ConcreteProductV2 具体产品
type ConcreteProductV2 struct {
}

// GetName 具体产品的实现方法
func (c *ConcreteProductV2) GetName() string {
	return "ConcreteProductV2"
}
