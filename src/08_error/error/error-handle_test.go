package err

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// *	在 golang 中 error 型態底層是 interface，
// *	這個 error 裡面有 Error() 這個 function，並且會回傳 string。
// *	只要任何 stuct 有實作 Error() 接口，就可以變成 error 物件
func TestErrors(t *testing.T) {
	err := errors.New("Customize error")
	fmt.Println(err.Error()) // Customize error
}

// Error Interface
// error 是 golang 內建的 interface，和 fmt.Stringer 很類似
type error interface {
	Error() string
}

// ----------------------------- 顯示錯誤訊息 -----------------------------------
// *	顯示錯誤訊息
// 		err.Error()
// 		fmt.Errorf("username %s is already existed", username)
func TestErrorMwssage(t *testing.T) {
	f, err := os.Open("filename.ext")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(f)
	}
}

// 產生錯誤訊息
// 作法：使用 errors.New()
// 直接回傳 error 的型別，並搭配 errors package 提供的 errors.New() 來產生錯誤訊息：
func checkUserNameExist(userName string) (bool, error) {
	if userName == "tom" {
		return true, errors.New("username tom is already exist")
	}
	return false, nil
}

func TestErrorsNew(t *testing.T) {
	if _, err := checkUserNameExist("tom"); err != nil {
		fmt.Println(err)  // username foo is already existed
	}
}

// 作法：使用 fmt.Errorf()
// 直接回傳 error 的型別，並搭配 fmt.Errorf() 來產生錯誤訊息：
func checkUserNameExist2(userName string) (bool, error) {
	if userName == "lin" {
		// 搭配 fmt.Errorf 來產生錯誤訊息
		return true, fmt.Errorf("userName %s is already existed", userName)
	}
	return false, nil
}

func TestFmtErrorf(t *testing.T) {
	if _, err := checkUserNameExist2("lin"); err != nil {
		fmt.Println(err) // userName lin is already existed
	}
}


// 	---------------------------- define error structure ----------------------------------
/*
		In many cases fmt.Errorf is good enough, but since error is an interface,
		you can use arbitrary data structures as error values,
		to allow callers to inspect the details of the error.
*/

// * Example 1:
// 自己定義 error structure（一）
type MyError struct {
	When time.Time
	What string
}

// STEP 2：定義能夠屬於 Error Interface 的方法,
// implement Error interface by defining method with MyError type Receiver
func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

// STEP 3：拋出錯誤的函式
func run() error {
	return MyError{
		When: time.Now(),
		What: "id didn't work",
	}
}

// STEP 4：使用 fmt.Println 即可取得錯誤拋出的訊息
func TestCustomizedError(t *testing.T) {
	if err := run(); err != nil {
		fmt.Println(err)
		// at 2022-05-21 16:20:10.59634 +0800 CST m=+0.000935185, id didn't work
	}
	err := MyError{time.Now(), "Err Message"}
	fmt.Println(err) // at 2022-05-21 19:28:46.708067 +0800 CST m=+0.001655977, Err Message
}

// *	另外，也可以使用 err.(ErrorStruct) 來判斷該 Error 是否屬於自己定義的錯誤型態：
// 		errorStruct, isErrorStruct := err.(ErrorStruct)
// STEP 1：定義另一個 Error struct
type NotMyError struct{}

func (f NotMyError) Error() string {
	return "foo"
}

func TestErrorType(t *testing.T) {
	if err := run(); err != nil {
		var isMyError bool
		// STEP 2：使用 err.(ErrorStruct) 來判斷該錯誤使否屬於該 Error Type
		_, isMyError = err.(NotMyError) // 判斷該錯誤是否屬於 NotMyError Type
		fmt.Println("isMyError: ", isMyError) // isMyError:  false

		_, isMyError = err.(MyError) 	// 判斷該錯誤是否屬於 MyError Type
		fmt.Println("isMyError: ", isMyError) // isMyError:  true
	}
}

// *	Example 2:
// 作法：自己定義 error structure（二）

// STEP 1：定義一個 error type
type ErrUserNameExist struct {
	UserName string
}

// STEP 2：擴充 Error == implement Error interface
func (e ErrUserNameExist) Error() string {
	return fmt.Sprintf("username %s is already exist", e.UserName)
}

func checkUserNameExist3(username string) (bool, error) {
	if username == "foo" {
		 // STEP 3：使用定義好的 error struct
		return true, ErrUserNameExist{
			UserName: username,
		}
	}
	return false, nil
}

func TestCustomizedError2(t *testing.T) {
	if _, err := checkUserNameExist3("foo"); err != nil {
		fmt.Println(err) // username foo is already exist
	}
}

// 也可以使用 err.(ErrorType) 進一步判斷這個 error type 是不是自己定義的 error structure（errUserNameExist）：
// STEP 1：判斷這個 err 是不是自己定義的 errUserNameExist
func isErrUserNameExist(err error) bool {
	// 因為 err 是 interface，裡面可以定義各種不同的形態
	_, ok := err.(ErrUserNameExist)
	return ok
}

func checkUserNameExist4(username string) (bool, error) {
	if username == "bar" {
		return true, errors.New("bar exist") // return errors.New("bar exist")
	}
	if username == "foo" {
		return true, ErrUserNameExist{ // return ErrUserNameExist implemented method
			UserName: username,
		}
	}
	return false, nil
}

func TestErrWithErrorTypeFunction(t *testing.T) {
	if _, err := checkUserNameExist4("foo"); err != nil {
		// STEP 2：進行判斷
		if isErrUserNameExist(err) {
			fmt.Println(err) // username foo is already exist
		}
	}
	if _, err := checkUserNameExist4("bar"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println(err) // bar exist
		}
	}
}

// ------------------------------ 錯誤發生時終止程式執行 ----------------------------------
// 錯誤發生時終止程式執行
// 若有需要在錯誤發生時終止程式執行，可以使用 os 提供的 os.Exit() 方法：
func TestErrorExit(t *testing.T) {
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	os.Exit(1)     // 終止程式繼續執行
	// }
	if _, err := checkUserNameExist4("foo"); err != nil {
		if isErrUserNameExist(err) {
			fmt.Println(err) // username foo is already exist
			os.Exit(1)
		}
	}
}

// ------------------------------ 判斷錯誤類型 ----------------------------------

// 判斷錯誤類型
// 使用 err.(*T) 的方式可以用來判斷錯誤的型別，例如 errMsg, isUnmarshalError := err.(*json.UnmarshalTypeError) ：

type SenserReading struct {
	Name string `json:"name"`
	Capacity int `json:"capacity"`
}

func TestErrorTypeWithJsonToStruct(t *testing.T) {
	jsonString := `{"name": "battery sensor", "capacity": "wrong time"}`
	reading := SenserReading{}

	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		unmarshalError, isUnmarshalError := err.(*json.UnmarshalTypeError) // 使用 err.(T) 來判斷錯誤的型別
		fmt.Println(unmarshalError) // json: cannot unmarshal string into Go struct field SensorReading.capacity of type int
		fmt.Println(isUnmarshalError) // true
	}
	fmt.Printf("%+v\n", reading) // {Name:battery sensor Capacity:0}
}

func TestErrorTypeWithJsonToStruct2(t *testing.T) {
	jsonString := `{"name": "battery sensor", "capacity": 5}`
	reading := SenserReading{}

	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		unmarshalError, isUnmarshalError := err.(*json.UnmarshalTypeError) // 使用 err.(T) 來判斷錯誤的型別
		fmt.Println(unmarshalError)
		fmt.Println(isUnmarshalError)
	}
	fmt.Printf("%+v\n", reading) // {Name:battery sensor Capacity:5}
}

// 也可以用來判斷客制的錯誤類型：
/*
func TestErrorType3(t *testing.T) {
	if pccErrors, ok := err.(*pcc.ErorResp); ok {
		firstError := pccErrors.Errors[0]
		statusCode, err := strconv.Atoi(firstError.Status)
		if err != nil {
			log.Warn("[api/helper] SuccessOrAbort - strconv.Atoi failed", err)
		}
	}
	_ = ctx.AbortWithError(statusCode, fmt.Errorf("%s:%s", firstError.Title, firstError.Detail))
  return
}
*/

// ----------------------------- Test Assertion -----------------------------------
func isEnable(enable bool) (bool, error) {
	if enable {
			return false, fmt.Errorf("You can't enable this setting")
	}

	return true, nil
}

func TestIsEnable(t *testing.T) {
	ok, err := isEnable(true)
	assert.False(t, ok)
	assert.NotNil(t, err)
	assert.Equal(t, "You can't enable this setting", err.Error())
}

