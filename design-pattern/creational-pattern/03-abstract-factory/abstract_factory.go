package _3_abstract_factory

// AbstractFactory 抽象工厂
type AbstractFactory interface {
	CreateProduct() AbstractProduct
}

// ConcreteFactoryV1 具体工厂
type ConcreteFactoryV1 struct {
}

// NewConcreteFactoryV1 创建具体工厂
func NewConcreteFactoryV1() *ConcreteFactoryV1 {
	return &ConcreteFactoryV1{}
}

// CreateProduct 创建产品
func (c *ConcreteFactoryV1) CreateProduct() AbstractProduct {
	return &ConcreteProductV1{}
}

// ConcreteFactoryV2 具体工厂
type ConcreteFactoryV2 struct {
}

// NewConcreteFactoryV2 创建具体工厂
func NewConcreteFactoryV2() *ConcreteFactoryV2 {
	return &ConcreteFactoryV2{}
}

// CreateProduct 创建产品
func (c *ConcreteFactoryV2) CreateProduct() AbstractProduct {
	return &ConcreteProductV2{}
}
