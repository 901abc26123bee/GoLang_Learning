package polymorph_test

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

type JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.println(\"Hello World\")"
}

// since Programmer is an interface, writeFirstProgram can only accept pointer of an instance as parameter
func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

func TestPolygon(t *testing.T) {
	goProg := &GoProgrammer{}
	javaProg := new(JavaProgrammer)
	writeFirstProgram(goProg)
	writeFirstProgram(javaProg)
}