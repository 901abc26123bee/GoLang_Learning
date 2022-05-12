package reflect

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func ChackType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
		case reflect.Float32, reflect.Float64:
			fmt.Println("Float")
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			fmt.Println("Int")
		default:
			fmt.Println("Unknown type", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	ChackType(f) // Float
	ChackType(&f) // Unknowntype *float64
}

func TestTypeAndValue(t *testing.T) {
	var f int64 = 10
	fmt.Print(reflect.TypeOf(f), reflect.ValueOf(f))
	fmt.Print(reflect.ValueOf(f).Type())
}

type Employee struct {
	EmployeeID string
	Name string `format:"normal"` // struct Tag
	Age int
}

type Customer struct {
	CookieID string
	Name string
	Age int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

// get properties in struct by name :
// 		reflect.ValueOf(*e).FieldByName("Name")
// get methods in struct by name :
// 		reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(20)})
// get properties' Tag in struct by name :
//		reflect.TypeOf(*e).FieldByName("Name") --> nameField.Tag.Get("format")
func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Yam", 30}
	fmt.Printf("Name: value(%[1]v), Type(%[1]T)", reflect.ValueOf(*e), reflect.TypeOf(*e))
	// Name: value({1 Yam 30}), Type(reflect.Value)Tag:format normal

	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("Failed to get Name Field")
	} else {
		fmt.Println("Tag:format", nameField.Tag.Get("format")) 
		// Tag:format normal
	}

	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(20)})
	fmt.Println("Updated Age: ", e) // Updated Age:  &{1 Yam 20}
}

// ----------------------------------------------------------------
//  map into aprcified strcut
func fillBySettings(st interface{}, settings map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() == reflect.Ptr {
		// st is a reference typr(pointer)
		// Elem[] --> get the value pointer point at --> get struct properties namd
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first parameter should be a pointer to a struct type")
		}
	}

	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok bool
	)

	for k, v := range settings {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}
	}
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Atom", "Age": 33}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	fmt.Println(e)

	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	fmt.Println(*c)
}