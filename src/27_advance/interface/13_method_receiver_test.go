package interface_test

import (
	"fmt"
	"testing"
)

type Cat2 struct {
	name           string
	scientificName string
	category       string
}

func New(name, scientificName, category string) Cat2 {
	return Cat2{
		name:           name,
		scientificName: scientificName,
		category:       category,
	}
}

func (cat *Cat2) SetName(name string) {
	cat.name = name
}

func (cat Cat2) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat2) Name() string {
	return cat.name
}

func (cat Cat2) ScientificName() string {
	return cat.scientificName
}

func (cat Cat2) Category() string {
	return cat.category
}

func (cat Cat2) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.category, cat.name)
}

func Test_Interface(t *testing.T) {
	cat := New("little pig", "American Shorthair", "cat")
	cat.SetName("monster") // (&cat).SetName("monster")
	fmt.Printf("The cat: %s\n", cat)

	cat.SetNameOfCopy("little pig")
	fmt.Printf("The cat: %s\n", cat)

	type Pet interface {
		SetName(name string)
		Name() string
		Category() string
		ScientificName() string
	}

	_, ok := interface{}(cat).(Pet)
	fmt.Printf("Cat2 implements interface Pet: %v\n", ok)
	_, ok = interface{}(&cat).(Pet)
	fmt.Printf("*Cat2 implements interface Pet: %v\n", ok)


	/*
		The cat: American Shorthair (category: cat, name: "monster")
		The cat: American Shorthair (category: cat, name: "monster")
		Cat2 implements interface Pet: false
		*Cat2 implements interface Pet: true
	*/
}