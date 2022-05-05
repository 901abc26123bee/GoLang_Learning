package map_test

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 5, 2: 6, 3: 7}
	t.Log(m1[2])
	t.Logf("len m1 = %d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m1 = %d", len(m2))

	m3 := make(map[int]int, 10) // make(type, capacity)
	t.Logf("len m1 = %d", len(m3))

}

func TestAccessNotExistsKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[2])

	m1[2] = 5
	t.Log(m1[2])

	m1[3] = 6

	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Logf("Key 3 is not exitsing")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	for k, v := range m1 {
		t.Logf(k, v)
	}
}