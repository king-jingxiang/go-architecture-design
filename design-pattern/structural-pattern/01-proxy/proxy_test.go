package _1_proxy

import "testing"

func TestProxy_Execute(t *testing.T) {
	proxy := NewProxy()
	r1 := proxy.Execute("Hi")
	if r1 != "Hi" {
		t.Errorf("Expected result to be 'Hi', got '%s'", r1)

	}
	r2 := proxy.Execute("yes")
	if r2 != "Proxy Service-yes" {
		t.Errorf("Expected result to be 'Proxy Service-yes', got '%s'", r2)

	}
}
