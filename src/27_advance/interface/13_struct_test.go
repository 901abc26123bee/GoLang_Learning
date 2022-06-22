package interface_test

import (
	"fmt"
	"testing"
)

type AnimalCategory struct {
  kingdom string // 界。
  phylum string // 門。
  class  string // 綱。
  order  string // 目。
  family string // 科。
  genus  string // 屬。
  species string // 種。
}
// 下邊有個名叫String的方法，從它的接收者聲明可以看出它隸屬於AnimalCategory類型。
// 通過該方法的接收者名稱ac，我們可以在其中引用到當前值的任何一個字段，或者調用到當前值的任何一個方法（也包括String方法自己）。
// 我們可以通過為一個類型編寫名為String的方法，來自定義該類型的字符串表示形式。
func (ac AnimalCategory) String() string {
  return fmt.Sprintf("%s%s%s%s%s%s%s",
  ac.kingdom, ac.phylum, ac.class, ac.order,
  ac.family, ac.genus, ac.species)
}
/*
in print.go

  Stringer is implemented by any value that has a String method,
  which defines the ``native'' format for that value.
  The String method is used to print values passed as an operand
  to any format that accepts a string or to an unformatted printer
  such as Print.

  type Stringer interface {
    String() string
  }
  ...
  func Sprintf(format string, a ...any) string {
    p := newPrinter()
    p.doPrintf(format, a)
    s := string(p.buf) // ***
    p.free()
    return s
  }
*/

func Test_Struct_Method(t *testing.T) {
  category := AnimalCategory{species: "cat"}
  fmt.Printf("The animal category: %s\n", category) // The animal category: cat
}

// ------------------------------------------------------------------------------------------------
type Animal struct {
  scientificName string // 學名。
  AnimalCategory    // 動物基本分類。
}

func (a Animal) Category() string {
  return a.AnimalCategory.String() // The animal: cat
}

func (a Animal) String() string {
  return fmt.Sprintf("%s (category: %s)",
  a.scientificName, a.AnimalCategory) // The animal: American Shorthair (category: cat)
}

func Test_Enbede_Field(t *testing.T) {
  category := AnimalCategory{species: "cat"}
  animal := Animal{
    scientificName: "American Shorthair",
    AnimalCategory: category,
  }
  // 只要名稱相同，無論這兩個方法的簽名是否一致，被嵌入類型的方法都會“屏蔽”掉嵌入字段的同名方法。
  // 不過，即使被屏蔽了，我們仍然可以通過鍊式的選擇表達式，選擇到嵌入字段的字段或方法
  fmt.Printf("The animal: %s\n", animal) // The animal: American Shorthair (category: cat)

}

// ------------------------------------------------------------------------------------------------
type Cat struct {
  name string
  Animal
}

func (cat Cat) String() string {
  return fmt.Sprintf("%s (category: %s, name: %q)",
  cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

func Test_Enbede_Field2(t *testing.T) {
  category := AnimalCategory{species: "cat"}
  animal := Animal{
    scientificName: "American Shorthair",
    AnimalCategory: category,
  }
  cat := Cat{
    name: "Meow~",
    Animal: animal,
  }
  // 只要名稱相同，無論這兩個方法的簽名是否一致，被嵌入類型的方法都會“屏蔽”掉嵌入字段的同名方法。
  // 不過，即使被屏蔽了，我們仍然可以通過鍊式的選擇表達式，選擇到嵌入字段的字段或方法
  fmt.Printf("The animal: %s\n", cat) // The animal: American Shorthair (category: cat, name: "Meow~")
}