package design_pattern

import "fmt"

//抽象工廠，可以產出多個產品對象

type Shape interface {
	Draw()
}

type Color interface {
	Fill()
}

type Circle struct {
}

type Square struct {
}

func (Circle) Draw() {
	fmt.Println("畫圓")
}

func (Square) Draw() {
	fmt.Println("畫方")
}

//實現紅色類別
type Red struct {
}

type Green struct {
}

func (Red) Fill() {
	fmt.Println("填充紅色")
}

func (Green) Fill() {
	fmt.Println("填充綠色")
}

//抽象工廠接口
type AbstractFactory interface {
	GetShape(shapeName string) Shape
	GetColor(shapeName string) Color
}

type ShapeFactory struct {
}

type ColorFactory struct {
}

func (sh ShapeFactory) GetShape(shapeName string) Shape {
	switch shapeName {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		return nil
	}
}

func (sh ShapeFactory) GetColor(shapeName string) Color {
	return nil
}

func (cf ColorFactory) GetShape(shapeName string) Shape {
	return nil
}

func (cf ColorFactory) GetColor(colorName string) Color {
	switch colorName {
	case "red":
		return &Red{}
	case "green":
		return &Green{}
	default:
		return nil
	}
}

type FactoryProducer struct {
}

func (fp FactoryProducer) GetFactory(factoryName string) AbstractFactory {
	switch factoryName {
	case "color":
		return &ColorFactory{}
	case "shape":
		return &ShapeFactory{}
	default:
		return nil
	}
}

func NewFactoryProducer() *FactoryProducer {
	return &FactoryProducer{}
}
