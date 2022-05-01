package extension_test

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("Pet ...")
}

func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

type Dog struct {
	p *Pet
}

func (d *Dog) Speak() {
	// using pet
	d.p.Speak()
	// self defined bt Dog
	// fmt.Print("Woof")
}

func (d *Dog) SpeakTo(host string) {
	// using pet
	d.p.SpeakTo(host)
	// self defined bt Dog
	// d.Speak()
	// fmt.Println(" ", host)
}

func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.SpeakTo("Lisa")
}

// Embedding interfaces
type Cat struct {
	Pet
}

// cannot override interface
func (d *Cat) Speak() {
	fmt.Print("Mow ~")
}

func TestCat(t *testing.T) {
	// var cat Pet = Pet(new(Cat)) // compoile error
	cat := new(Cat)
	cat.SpeakTo("Tom")
}
