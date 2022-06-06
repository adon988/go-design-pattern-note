package main

import (
	"fmt"
	"myproject/design_pattern"
)

func main() {

	//factory
	apple := design_pattern.NewFruit(1)
	apple.Grant()
	apple.Pick()

	orange := design_pattern.NewFruit(2)
	orange.Grant()
	orange.Pick()

	//singleton
	s1 := design_pattern.GetInstance()
	fmt.Printf("s1: %p\n", s1)
	s2 := design_pattern.GetInstance()
	fmt.Printf("s2: %p\n", s2)
	s2.Dosomething()

	//Abstract Factory
	superFactory := design_pattern.NewFactoryProducer()
	colorFactory := superFactory.GetFactory("color")
	shapeFactory := superFactory.GetFactory("shape")
	red := colorFactory.GetColor("red")
	green := colorFactory.GetColor("green")

	circle := shapeFactory.GetShape("circle")
	square := shapeFactory.GetShape("square")

	circle.Draw()
	red.Fill()
	green.Fill()
	square.Draw()

	//Builder
	lc := design_pattern.LowComputer{}
	d2 := design_pattern.NewConstructor(&lc)
	d2.Construct()

	hc := design_pattern.SuperComputer{}
	hd := design_pattern.NewConstructor(&hc)
	hd.Construct()
}
