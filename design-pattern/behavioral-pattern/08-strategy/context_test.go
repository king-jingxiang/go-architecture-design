package _8_strategy

import "testing"

func TestContext_Execute(t *testing.T) {
	strategyA := NewStrategyA()
	strategyB := NewStrategyB()
	context := NewContext()
	context.SetStrategy(strategyA)
	r1 := context.Execute()
	if r1 != "Strategy A" {
		t.Errorf("Expected result to be 'Strategy A', got '%s'", r1)

	}
	context.SetStrategy(strategyB)
	r2 := context.Execute()
	if r2 != "Strategy B" {
		t.Errorf("Expected result to be 'Strategy B', got '%s'", r2)

	}
}
