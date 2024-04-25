package _3_abstract_factory

import (
	"testing"
)

func TestAbstractFactory_CreateProduct(t *testing.T) {
	factory := NewConcreteFactoryV1()
	product := factory.CreateProduct()
	if product.GetName() != "ConcreteProductV1" {
		t.Errorf("Expected product name to be 'ConcreteProductV1', got '%s'", product.GetName())
	}
	factoryV2 := NewConcreteFactoryV2()
	product = factoryV2.CreateProduct()
	if product.GetName() != "ConcreteProductV2" {
		t.Errorf("Expected product name to be 'ConcreteProductV2', got '%s'", product.GetName())
	}
}
