package design_pattern

import (
	"fmt"
)

//工廠模式，一個具體工廠產生一個具體產品

//種植及採集水果
type Fruit interface {
	Grant()
	Pick()
}

type Apple struct {
	name string
}

type Orange struct {
	name string
}

func (*Apple) Grant() {
	fmt.Println("種植蘋果")
}

func (*Apple) Pick() {
	fmt.Println("採集蘋果")
}

func (*Orange) Pick() {
	fmt.Println("採集橘子")
}

func (*Orange) Grant() {
	fmt.Println("種植橘子")
}

func NewFruit(t int) Fruit {
	switch t {
	case 1:
		return &Apple{}
	case 2:
		return &Orange{}
	default:
		return nil
	}
}
