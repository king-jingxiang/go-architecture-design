package _1_singleton

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	// 测试GetInstance是否总是返回同一个实例
	inst1 := GetInstance()
	inst2 := GetInstance()
	if inst1 != inst2 {
		t.Errorf("Expected same instance but got different instances")
	}
}

func TestGetInstanceConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	instances := make([]*singleton, 0)
	lock := sync.Mutex{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inst := GetInstance()
			lock.Lock()
			defer lock.Unlock()
			instances = append(instances, inst)
		}()
	}
	wg.Wait()

	firstInstance := instances[0]
	for _, inst := range instances[1:] {
		if inst != firstInstance {
			t.Errorf("Expected same instance in all goroutines but got different instances")
		}
	}
}
