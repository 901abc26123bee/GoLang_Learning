package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEquals(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	c := map[int]string{1: "one", 3: "three", 2: "two"}
	// fmt.Println(a == b) --> reference can not use "==" tp compare
	fmt.Println(reflect.DeepEqual(a, b)) // true
	fmt.Println(reflect.DeepEqual(a, c)) // true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 1, 3}
	fmt.Println("s1 == s2 ? ", reflect.DeepEqual(s1, s2)) // s1 == s2 ?  true
	fmt.Println("s1 == s3 ? ", reflect.DeepEqual(s1, s3)) // s1 == s3 ?  false
}