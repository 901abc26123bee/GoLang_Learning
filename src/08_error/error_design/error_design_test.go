package errordesign_test

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"testing"
)

type Point struct {
	Longitude float64
	Latitude float64
	Distance float64
	ElevationGain float64
	ElevationLoss float64
}

// https://coolshell.cn/articles/21140.html
// ------------------------------------------------------------------------------------------------
// Error Check  Hell
func parse(r io.Reader) (*Point, error) {
	var p Point
	if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
			return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
			return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
			return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
			return nil, err
	}
	if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
			return nil, err
	}
	return &p, nil
}

// ------------------------------------------------------------------------------------------------
// using functional programming
// 我們通過使用Closure 的方式把相同的代碼給抽出來重新定義一個函數，這樣大量的  if err!=nil 處理的很乾淨了。
// 但是會帶來一個問題，那就是有一個 err 變量和一個內部的函數，感覺不是很乾淨
func parse_better(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
			if err != nil {
					return
			}
			err = binary.Read(r, binary.BigEndian, data)
	}
	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)
	if err != nil {
			return &p, err
	}
	return &p, nil
}

// ------------------------------------------------------------------------------------------------
// 我們從Go 語言的 bufio.Scanner()中似乎可以學習到一些東西
// scanner在操作底層的I/O的時候，那個for-loop中沒有任何的 if err !=nil 的情況，
// 退出循環後有一個 scanner.Err() 的檢查。看來使用了結構體的方式。模仿它，我們可以把我們的代碼重
func FromBufioScannerSourceCode() {
	/*
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
			token := scanner.Text()
			// process token
	}
	if err := scanner.Err(); err != nil {
			// process the error
	}
	*/
}


// 首先，定義一個結構體和一個成員函數
type Reader struct {
	r   io.Reader
	err error
}
func (r *Reader) read(data interface{}) {
	if r.err == nil {
			r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

// 然後，我們的代碼就可以變成下面這樣
func parse3(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}
	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)
	if r.err != nil {
			return nil, r.err
	}
	return &p, nil
}

// ------------------------------------------------------------------------------------------------
// 有了上面這個技術，我們的“流式接口 Fluent Interface”，也就很容易處理了。如下所示：

// 長度不夠，少一個Weight --> leed to error
var b = []byte {0x48, 0x61, 0x6f, 0x20, 0x43, 0x68, 0x65, 0x6e, 0x00, 0x00, 0x2c}
var r = bytes.NewReader(b)
type Person struct {
  Name [10]byte
  Age uint8
  Weight uint8
  err error
}
func (p *Person) read(data interface{}) {
  if p.err == nil {
    p.err = binary.Read(r, binary.BigEndian, data)
  }
}
func (p *Person) ReadName() *Person {
  p.read(&p.Name)
  return p
}
func (p *Person) ReadAge() *Person {
  p.read(&p.Age)
  return p
}
func (p *Person) ReadWeight() *Person {
  p.read(&p.Weight)
  return p
}
func (p *Person) Print() *Person {
  if p.err == nil {
    fmt.Printf("Name=%s, Age=%d, Weight=%d\n",p.Name, p.Age, p.Weight)
  }
  return p
}
func TestFunctionalStream(t *testing.T) {
	// []byte to string
  //  s2 := string(b)
	/*
  string to []byte
	s1 := "dora"
	b := []byte(s1)
	b10 := [10]byte{1,0,0,1,1,1,1,1,4,9}
	fmt.Println(b, b10)
	p := Person{b10, 22, 40, errors.New("Customize error")}
	*/
	p := Person{}
  p.ReadName().ReadAge().ReadWeight().Print()
  fmt.Println(p.err)  // EOF 错误
}

// ------------------------------------------------------------------------------------------------
// 包裝錯誤
// 最後，多說一句，我們需要包裝一下錯誤，而不是乾巴巴地把err給返回到上層，我們需要把一些執行的上下文加入。

// 通常来说，我们会使用 fmt.Errorf()来完成这个事，比如：

/*
	if err != nil {
		return fmt.Errorf("something failed: %v", err)
	}
*/


// 另外，在Go語言的開發者中，更為普遍的做法是將錯誤包裝在另一個錯誤中，同時保留原始內容：
type authorizationError struct {
	operation string
	err error   // original error
}

func (e *authorizationError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}

// 當然，更好的方式是通過一種標準的訪問方法，
// 這樣，我們最好使用一個接口，比如 causer接口中實現 Cause() 方法來暴露原始錯誤，以供進一步檢查：
type causer interface {
	Cause() error
}

func (e *authorizationError) Cause() error {
	return e.err
}


// 有一個第三方的錯誤庫（github.com/pkg/errors），go get github.com/pkg/errors
// 對於這個庫，我無論到哪都能看到他的存在，所以，這個基本上來說就是事實上的標準了。代碼示例如下：
/*
func TestErrorsPkg(t *testing.T) {
	//*錯誤包裝
	if err != nil {
		return errors.Wrap(err, "read failed")
	}
	//*Cause接口
	switch err := errors.Cause(err).(type) {
	case *MyError:
		//*handle specifically
	default:
		//*unknown error
	}
}
*/
