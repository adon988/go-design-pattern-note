package design_pattern

import "fmt"

type Builder interface {
	BuildDisc()
	BuildCPU()
	BuildRom()
}

type SuperComputer struct {
	Name string
}

func (this *SuperComputer) BuildDisc() {
	fmt.Println("Building super Disc")
}

func (this *SuperComputer) BuildCPU() {
	fmt.Println("Building CPU")
}

func (this *SuperComputer) BuildRom() {
	fmt.Println("Building ROM")
}

type LowComputer struct {
	Name string
}

func (this *LowComputer) BuildDisc() {
	fmt.Println("Building Low Disc")
}

func (this *LowComputer) BuildCPU() {
	fmt.Println("Building Low CPU")
}

func (this *LowComputer) BuildRom() {
	fmt.Println("Building Low ROM")
}

type Dictor struct {
	build Builder
}

func NewConstructor(b Builder) *Dictor {
	return &Dictor{
		build: b,
	}
}

func (this *Dictor) Construct() {
	this.build.BuildDisc()
	this.build.BuildCPU()
	this.build.BuildRom()
}
