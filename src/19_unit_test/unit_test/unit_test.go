package unittest

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestSquares(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}

	for i := 0; i < len(inputs); i++ {
		ret :=squares(inputs[i])
		assert.Equal(t, expected[i], ret)
		// if ret!= expected[i] {
		// 	t.Errorf("Input us %d, the expected is %d, the actual is %d", inputs[i], expected[i], ret)
		// }
	}
}

func TestErrorInCode(t *testing.T) {
	fmt.Println("Start")
	t.Error("Error")
	fmt.Println("End")
	// Start
	// Error
	// End
}

func TestFailInCode(t *testing.T) {
	fmt.Println("Start")
	t.Fatal("Error")
	fmt.Println("End")
	// Start
	// Error
}

// test coverage : go test -v -cover
// ~/Desktop/go_learning/src/19_unit_test/unit_test $ go test -v -cover
/*
=== RUN   TestSquares
    unit_test.go:16: Input us 1, the expected is 1, the actual is 2
    unit_test.go:16: Input us 2, the expected is 4, the actual is 5
    unit_test.go:16: Input us 3, the expected is 9, the actual is 10
--- FAIL: TestSquares (0.00s)
=== RUN   TestErrorInCode
Start
    unit_test.go:23: Error
End
--- FAIL: TestErrorInCode (0.00s)[]
=== RUN   TestFailInCode
Start
    unit_test.go:32: Error
--- FAIL: TestFailInCode (0.00s)
FAIL
coverage: 100.0% of statements
exit status 1
*/


// test assertion : https//github.com/stretchr/testify
// $ go get -u github.com/stretchr/testify/assert