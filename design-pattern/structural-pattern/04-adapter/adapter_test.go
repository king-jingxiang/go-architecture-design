package _4_adapter

import "testing"

func TestObjectAdapter_Execute(t *testing.T) {
	adapter := &ObjectAdapterV1{}
	result := adapter.Execute()
	if result != "Final Execute Result V1" {
		t.Errorf("Expected result to be 'Final Execute Result V1', got '%s'", result)
	}
	adapter2 := &ObjectAdapterV2{}
	result = adapter2.Execute()
	if result != "Final Execute Result V2" {
		t.Errorf("Expected result to be 'Final Execute Result V2', got '%s'", result)
	}
}
