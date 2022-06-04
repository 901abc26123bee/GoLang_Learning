package reversecontrol

import "errors"

// [GO编程模式：委托和反转控制](https://coolshell.cn/articles/21214.html)

/*
type IntSetOld struct {
	data map[int]bool
}

// 其中實現了 Add() 、Delete() 和 Contains() 三個操作，前兩個是寫操作，後一個是讀操作。
func NewIntSet() IntSet {
	return IntSet{make(map[int]bool)}
}
func (set *IntSet) Add(x int) {
	set.data[x] = true
}
func (set *IntSet) Delete(x int) {
	delete(set.data, x)
}
func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}
*/

// 現在我們想實現一個 Undo 的功能。我們可以把把 IntSet 再包裝一下變成 UndoableIntSet
type UndoableIntSet struct {
	IntSet // Embedding (delegation)
	functions []func()
}

func NewUndoableIntSet() UndoableIntSet {
	return UndoableIntSet{NewIntSet(), nil}
}

// 我們在 UndoableIntSet 中嵌入了IntSet ，然後Override了 它的 Add()和 Delete() 方法。
// Contains() 方法沒有Override，所以，會被帶到 UndoableInSet 中來了。
// 在Override的 Add()中，記錄 Delete 操作
// 在Override的 Delete() 中，記錄 Add 操作
// 新加入 Undo() 中進行Undo操作。
// * Undo操作其實是一種控制邏輯，並不是業務邏輯，所以，在復用 Undo這個功能上是有問題。因為其中加入了大量跟 IntSet 相關的業務邏輯。

func (set *UndoableIntSet) Add(x int) {  // Override
	if !set.Contains(x) {
		set.data[x] = true
		set.functions = append(set.functions, func() {set.Delete(x)})
	} else {
		set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Delete(x int) { // Override
	if set.Contains(x) {
			delete(set.data, x)
			set.functions = append(set.functions, func() { set.Add(x) })
	} else {
			set.functions = append(set.functions, nil)
	}
}

func (set *UndoableIntSet) Undo() error {
	if len(set.functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(set.functions) - 1
	if function := set.functions[index]; function != nil {
		function()
		set.functions[index] = nil // For garbage collection
	}
	set.functions = set.functions[:index]
	return nil
}

// 反轉依賴
// 我們先聲明一種函數接口，表現我們的Undo控制可以接受的函數簽名是什麼樣的：
type Undo []func()

// 有了上面這個協議後，我們的Undo控制邏輯就可以寫成如下：
func (undo *Undo) Add(function func()) {
  *undo = append(*undo, function)
}
func (undo *Undo) Undo() error {
  functions := *undo
  if len(functions) == 0 {
    return errors.New("No functions to undo")
  }
  index := len(functions) - 1
  if function := functions[index]; function != nil {
    function()
    functions[index] = nil // For garbage collection
  }
  *undo = functions[:index]
  return nil
}

// 然後，我們在我們的IntSet裡嵌入 Undo，然後，再在 Add() 和 Delete() 裡使用上面的方法，就可以完成功能。
type IntSet struct {
	data map[int]bool
	undo Undo
}

func NewIntSet() IntSet {
	return IntSet{data: make(map[int]bool)}
}
func (set *IntSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntSet) Contains(x int) bool {
	return set.data[x]
}
func (set *IntSet) Add(x int) {
	if !set.Contains(x) {
			set.data[x] = true
			set.undo.Add(func() { set.Delete(x) })
	} else {
			set.undo.Add(nil)
	}
}

func (set *IntSet) Delete(x int) {
	if set.Contains(x) {
			delete(set.data, x)
			set.undo.Add(func() { set.Add(x) })
	} else {
			set.undo.Add(nil)
	}
}

// 這個就是控制反轉，不再由 控制邏輯 Undo 來依賴業務邏輯 IntSet，而是由業務邏輯 IntSet 來依賴 Undo 。
// 其依賴的是其實是一個協議，這個協議是一個沒有參數的函數數組。我們也可以看到，我們 Undo 的代碼就可以復用了。