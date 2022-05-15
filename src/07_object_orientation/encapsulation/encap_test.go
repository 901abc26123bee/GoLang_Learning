package encapsulation_test

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}

func TestCreateEmployee(t *testing.T) {
	e := Employee{"0", "Kin", 20}
	e1 := Employee{Name: "John", Age:30}
	e2 := new(Employee)
	e2.Name = "Atom"
	e2.Age = 40
	e2.Id = "3"

	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Log(e2.Name)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e) // pointer
	t.Logf("e is %T", e2) // pointer
}

// avoid copy in memory
// func (e *Employee) String() string {
// 	fmt.Printf("Adress is %x ", unsafe.Pointer(&e.Name))
// 	return fmt.Sprintf("ID: %s-Name: %s Age: %d", e.Id, e.Name, e.Age)
// }

// Employee e passed by copy in memory
func (e Employee) String() string {
	fmt.Printf("Adress is %x ", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID: %s-Name: %s Age: %d", e.Id, e.Name, e.Age)
}

func TestStructOperation(t *testing.T) {
	// e2 := &Employee{"0", "Jen", 22}
	e := Employee{"0", "Jen", 22}
	fmt.Printf("Adress is %x ", unsafe.Pointer(&e.Name))
	t.Logf(e.String())
}