package _1_singleton

import "sync"

var instance *singleton
var once sync.Once

type singleton struct {
	name string
}

func GetInstance() *singleton {
	once.Do(
		func() {
			instance = &singleton{
				name: "singleton",
			}
		},
	)
	return instance
}
