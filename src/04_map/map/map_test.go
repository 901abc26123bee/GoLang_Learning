package map_test

import (
	"fmt"
	"testing"
)

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

// 建立 Map 的方式
// 在 Go 中，Map 也是 Key-Value pair 的組合，
// 但是 Map 所有 Key 的資料型別都要一樣；所有 Value 的資料型別也要一樣。
// 另外，在設值和取值需要使用 [] 而非 .。map 的 zero value 是 nil。
func TestMapCreation(t *testing.T) {
  // *第一種方式
  // 建立一個型別為 map 的變數，其中的 key 都會是 string，value 也都會是 string
  colors := map[string]string{}
	fmt.Println(colors) // map[]

  // 也可以直接帶值
	colors2 := map[string]string{
			"red":   "#ff0000",
			"green": "#4bf745",
	}
	fmt.Println(colors2) // map[green:#4bf745 red:#ff0000]

  // *第二種方式
  var colors3 map[string]string
	fmt.Println(colors3) // map[]

  // *第三種方式，使用 make 建立 Map。
  // Key 的型別是 string，Value 是 int
  colors4 := make(map[string]int)
	colors4["red"] = 10
	fmt.Println(colors4) // map[red:10]

  // 如果對於鍵值的數量有概念的話，也可以給定初始大小，將有助於效能提升（The little Go book）
  colors5 := make(map[string]int, 100)
	fmt.Println(colors5) // map[]
}

// Map 的 value 是 struct
type Vertex struct {
	Lat, Long float64
}
// smae as Vertex2
type Vertex2 struct {
	Lat float64
	Long float64
}


func TestMapValueIsStruct(t *testing.T) {
// 使用 make 建立 Map
	m := make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
			40.68443, -74.39967,
	}


	// 使用 Map Literal 建立 Map
	mapLiteral := map[string]Vertex{
			"Bell Labs": Vertex{
					40.68433, -74.39967,
			},
			"Google": Vertex{
					37.42202, -122.08408,
			},
	}

	// Struct Type 的名稱可以省略
	mapLiteral2 := map[string]Vertex{
			"Bell Labs": {
					40.68433, -74.39967,
			},
			"Google": {
					37.42202, -122.08408,
			},
	}
	
	fmt.Println(mapLiteral)
	fmt.Println(mapLiteral2)
}