package Singleton

import "sync"

var checkOnce = sync.Once{}

type DpSingleton struct {
	data string
}

var instance *DpSingleton

// GetInstance handle to initialize the instance just once
func GetInstance(newData string) *DpSingleton {
	if instance == nil {
		checkOnce.Do(func() {
			if instance == nil {
				instance = &DpSingleton{data: newData}
			}
		})
	}
	return instance
}
