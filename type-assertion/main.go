package main

import "fmt"

type Plant interface {
	Photosynthesise()
}

type Rose struct{}

func (Rose) Photosynthesise() {}

type Orchid struct{}

func (Orchid) Photosynthesise() {}

type Human struct{}

func (Human) Metabolise() {}

type Sun struct{}

func (Sun) feed(Plant) {}

func main() {
	// plants can photosynthesise
	Sun{}.feed(Rose{})

	// humans cannot photosynthesise
	// Sun{}.feed(Human{})

	// not allowed to cast non-interface type
	// var s string
	// _ = s.(int)

	// compiler/linter does not warn about silly type assertion (at least mine doesn't),
	var p Plant = Rose{}
	// _ = p.(Orchid) // this panics
	o, ok := p.(Orchid) // this doesn't panic
	fmt.Println(o, ok)  // {} false
}
