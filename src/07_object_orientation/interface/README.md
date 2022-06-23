https://github.com/aceld/golang/blob/main/4%E3%80%81interface.md
https://www.henrydu.com/2021/02/07/golang-interface-brief-look/
https://draveness.me/golang/docs/part2-foundation/ch04-basic/golang-interface/

interface 是GO語言的基礎特性之一。可以理解為一種類型的規範或者約定。它是通過約定的形式，隱式的實現interface 中的方法即可。因此，Golang 中的 interface 讓編碼更靈活、易擴展。
1. interface 是方法聲明的集合
2. 任何類型的對象實現了在interface 接口中聲明的全部方法，則表明該類型實現了該接口。
3. interface 可以作為一種數據類型，實現了該接口的任何對像都可以給對應的接口類型變量賦值。
interface 可以被任意對象實現，一個類型/對像也可以實現多個 interface 　　b. 方法不能重載，如 eat(), eat(s string) 不能同時存在


interface在使用的過程中，共有兩種表現形式 一種為空接口(empty interface)，定義如下：
  ```go
    var MyInterface interface{}
  ```
另一種為非空接口(non-empty interface), 定義如下：
  ```go
    type MyInterface interface { function() }
  ```
這兩種interface類型分別用兩種struct表示，空接口為`eface`, 非空接口為`iface`.

### 空接口eface
空接口`eface`結構，由兩個屬性構成，一個是類型信息`_type`，一個是數據信息。其數據結構聲明如下：
  ```go
  type eface struct {      //空接口
      _type *_type         //類型信息
      data  unsafe.Pointer //指向數據的指針(go語言中特殊的指針類型unsafe.Pointer類似於c語言中的void*)
  }
  ```

_type屬性：是GO語言中所有類型的公共描述，Go語言幾乎所有的數據結構都可以抽象成 _type，是所有類型的公共描述，**type負責決定data應該如何解釋和操作，**type的結構代碼如下:
  ```go
  type _type struct {
      size       uintptr  //類型大小
      ptrdata    uintptr  //前綴持有所有指針的內存大小
      hash       uint32   //數據hash值
      tflag      tflag
      align      uint8    //對齊
      fieldalign uint8    //嵌入結構體時的對齊
      kind       uint8    //kind 有些枚舉值kind等於0是無效的
      alg        *typeAlg //函數指針數組，類型實現的所有方法
      gcdata    *byte
      str       nameOff
      ptrToThis typeOff
  }
  ```
data屬性: 表示指向具體的實例數據的指針，他是一個unsafe.Pointer類型，相當於一個C的萬能指針void*。

### 非空接口iface
iface 表示 non-empty interface 的數據結構，非空接口初始化的過程就是初始化一個iface類型的結構，其中data的作用同eface的相同，這裡不再多加描述。
```go
type iface struct {
  tab  *itab
  data unsafe.Pointer
}
```
iface結構中最重要的是itab結構（結構如下），每一個 itab 都佔 32 字節的空間。 itab可以理解為pair<interface type, concrete type> 。 itab裡麵包含了interface的一些關鍵信息，比如method的具體實現。
```go
type itab struct {
  inter  *interfacetype   // 接口自身的元信息
  _type  *_type           // 具體類型的元信息
  link   *itab
  bad    int32
  hash   int32            // _type裡也有一個同樣的hash，此處多放一個是為了方便運行接口斷言
  fun    [1]uintptr       // 函數指針，指向具體類型所實現的方法
}
```

```go
type itab struct {
	inter *interfacetype
	_type *_type
	hash  uint32
	_     [4]byte
	fun   [1]uintptr // the pointer to point the set of methods.
}

type interfacetype struct {
	typ     _type
	pkgpath name
	mhdr    []imethod
}
```
- interfacetype包含了一些關於interface本身的信息，比如package path，包含的method。這裡的interfacetype是定義interface的一種抽象表示。
- type表示具體化的類型，與eface的 type類型相同。
- hash字段其實是對_type.hash的拷貝，它會在interface的實例化時，用於快速判斷目標類型和接口中的類型是否一致。另，Go的interface的Duck-typing機制也是依賴這個字段來實現。
- fun字段其實是一個動態大小的數組，雖然聲明時是固定大小為1，但在使用時會直接通過fun指針獲取其中的數據，並且不會檢查數組的邊界，所以該數組中保存的元素數量是不確定的。
