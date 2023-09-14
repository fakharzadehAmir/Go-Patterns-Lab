package Singleton

import (
	"sync"
	"testing"
)

func TestSingleton(t *testing.T) {
	instance1 := GetInstance("HelloInstance1")
	instance2 := GetInstance("HelloInstance2")
	//	Ensure that two instances are the same
	if instance1 != instance2 {
		t.Errorf("Expected both instance1 and instance2 to be the same.")
	}
	//	Ensure that the data is set correctly
	expectedData := "HelloInstance1"
	if instance1.data != expectedData && instance2.data != expectedData {
		t.Errorf("Expected data to be %s, but got %s", instance1.data, expectedData)
	}
}

func TestSingletonConcurrent(t *testing.T) {
	//	Test concurrent
	numberOfGoroutines := 100000000
	instances := make([]*DpSingleton, numberOfGoroutines)
	var wg sync.WaitGroup
	for i := 0; i < numberOfGoroutines; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			instances[index] = GetInstance("HelloInstance" + string(rune(index+1)))
		}(i)
	}
	wg.Wait()

	//	Ensure that all instances are the same
	for i := 1; i < numberOfGoroutines; i++ {
		if instances[i] != instances[i-1] {
			t.Errorf("Expected all instances to be the same.")
		}
	}
}
