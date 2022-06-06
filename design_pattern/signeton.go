package design_pattern

import (
	"fmt"
	"sync"
)

//使用 sync.Once 來實現單例模式

type Singleton interface {
	Dosomething()
}

type SingletonSt struct {
}

func (*SingletonSt) Dosomething() {
	fmt.Println("do some thing in singleton")
}

var (
	once     sync.Once
	instance Singleton
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &SingletonSt{}
	})
	return instance
}
