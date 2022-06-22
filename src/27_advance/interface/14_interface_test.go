package interface_test

import (
	"testing"
	"fmt"
)

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func Test_Interface_2(t *testing.T) {
	// 如果我們使用一個變量給另外一個變量賦值，那麼真正賦給後者的，並不是前者持有的那個值，而是該值的一個副本。
	dog := Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	var pet Pet = dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())
	fmt.Println()

	dog1 := Dog{"little pig"}
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	dog2 := dog1
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	dog1.name = "monster"
	fmt.Printf("The name of first dog is %q.\n", dog1.Name())
	fmt.Printf("The name of second dog is %q.\n", dog2.Name())
	fmt.Println()

	dog = Dog{"little pig"}
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	pet = &dog
	dog.SetName("monster")
	fmt.Printf("The dog's name is %q.\n", dog.Name())
	fmt.Printf("This pet is a %s, the name is %q.\n",
		pet.Category(), pet.Name())

	/*
		The dog's name is "little pig".
		The dog's name is "monster".
		This pet is a dog, the name is "little pig".

		The name of first dog is "little pig".
		The name of second dog is "little pig".
		The name of first dog is "monster".
		The name of second dog is "little pig".

		The dog's name is "little pig".
		The dog's name is "monster".
		This pet is a dog, the name is "monster".
	*/
}