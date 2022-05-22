package json_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

// json.Unmarshal 把 json 字串轉成 struct
// json.Marshal 把 struct 轉成 json 字串

// ------------------------------------------------------------------------------------------------
// *	將 struct 轉成 JSON
// 		定義可轉換成 JSON 的 struct
// 		透過 struct field's tag 可以將 struct 轉換成特定的格式：

// struct fields tag: `json:"<field>,<options>"`
type Book struct {
	BookID        uint   `json:"-"`              // 轉換時總是忽略掉此欄位
	Title         string `json:"title"`          // 轉換後 JSON 的欄位名稱是 title
	Author        string `json:"author"`         // 轉換後 JSON 的欄位名稱是 author
	Name          string `json:"name,omitempty"` // 當 name 是空值時轉換後則無該欄位
	Age           uint   `json:",omitempty"`     // 當 Age 有值時，則 JSON 的欄位名稱是 "Age"（大寫開頭），否則不顯示該欄位
	Price         uint   `json:"_,"`             // 轉換後 JSON 的欄位名稱是 _
	Configuration string `json:"configuration,string"`
}

/*
	-：field tag 如果是 - 則總是忽略掉該欄位，轉換後不會有該欄位
	在 options 可以有一寫特殊的選項：
		omitempty ：如果該 field 是空值，則忽略該 field 的欄位，
								空值包含 false, 0, nil pointer, nil interface value, 空陣列、空字串
		string：如果某欄位的值又是 JSON 的話，可以在 options 中使用 string
*/

/*
	使用 Marshal 方法
	
	把 struct 放入 json.Marshal(struct) 的方法中，會回傳 byte slice
	透過 string(byteSlice) 把 byte slice 轉成 string 即可得到 JSON
	在將 struct 轉換成 JSON 時會根據不同的型別轉換出不同的內容，多數的轉換都相當直覺能夠直接對應（例如 int 變成 number），比較特別的是
	- []byte 會轉換成 base64-encoded 的字串
	- nil 會轉換成 null
*/
func TestStructToJson(t *testing.T) {
	book := Book{BookID: 2, Title: "Learning Go", Author: "Gopher", Name: "", Age: 0, Price: 31900}

	/* 將 struct 轉成 byte slice，再透過 string 變成 JSON 格式 */
	byteSlice, _ := json.MarshalIndent(book, "", "  ")
	fmt.Println(string(byteSlice))

	//{
	//  "title": "Learning Go",
	//  "author": "Gopher",
	//  "_": 31900,
	//  "configuration": "\"\""
	//}
}


//	*  將 struct 轉成 JSON（nested Object）
// Book struct 裡包含 Author struct
type Book2 struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

func TestNestedStructToJson(t *testing.T) {
	book1 := Book2{
		Title: "Learning Go",
		Author: Author{
			Sales: 3,
			Age: 25,
			Developer: true,
		},
	}

	author := Author{
		Sales:     3,
		Age:       25,
		Developer: true,
	}
	book2 := Book2{Title: "Learning Go", Author: author}

	/* 將 struct 轉成 byte slice，再透過 string 變成 JSON 格式 */
	byteSlice1, _ := json.MarshalIndent(book1, "", "  ")
	fmt.Println(string(byteSlice1))

	byteSlice2, _ := json.MarshalIndent(book2, "", "  ")
	fmt.Println(string(byteSlice2))
	/*
	{
		"title": "Learning Go",
		"author": {
			"book_sales": 3,
			"age": 25,
			"is_developer": true
		}
	}

	{
		"title": "Learning Go",
		"author": {
			"book_sales": 3,
			"age": 25,
			"is_developer": true
		}
	}
	*/
}

/*
 *	將 JSON 轉成 struct：結構化資料
		- 定義 JSON 檔案轉換後的 struct
		- 將 JSON 轉成 byte slice
		- 透過 json.Unmarshal(byte_slice, &struct) 將 byte slice 轉成 struct
		⚠️ 如果要接收來自 JavaScript 時間格式，建議傳送 ISO 8601 的格式（而非時間戳記）會比較簡單，
			也就是 new Date().toISOString()，如果傳送的是 timestamp 可能會需要另外處理。
			Format a time or date [complete guide](https://yourbasic.org/golang/format-parse-string-time-date-example/)
*/
type SensorReading struct {
	Name 			string	 	`json:"name"`
	Capacity  int 			`json:"capacity"`
	Time 			time.Time `json:"time"`
}

func TestJsonToStruct(t *testing.T) {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	reading := SensorReading{}

	/**
	* 將 JSON 轉成 struct 需要先把 string 轉成 byte slice，
	* 然後再透過 Unmarshal 把空的 Struct 帶入
	**/
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading) 
	// {Name:battery sensor Capacity:40 Time:2019-01-21 19:07:28 +0000 UTC}
}

// *	將 JSON 轉成 map：非結構化資料（unstructured data）
// *	有些時候並沒有辦法事先知道 JSON 資料的格式會長什麼樣子，
// *	這時候可以使用 map[string]interface{} 這種寫法，如此將可以把 JSON 轉成 map：
func TestJsonToMap(t *testing.T) {
	jsonString := `{"name": "battery sensor", "capacity": 40, "time": "2019-01-21T19:07:28Z"}`
	reading := make(map[string]interface{})

	/* 將 JSON 轉成 struct 需要先把 string 轉成 byte slice，然後再透過 Unmarshal 把空的 Struct 帶入 */
	err := json.Unmarshal([]byte(jsonString), &reading)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", reading)
	// map[capacity:40 name:battery sensor time:2019-01-21T19:07:28Z]
}

// *	將 JSON 轉成 struct：將 API 取得的 response 進行 JSON decode
// 		keywords: json.NewDecoder, Decode
// 將 resp.Body 傳入 json.NewDecoder() 中，在使用 Decode 來將 JSON 轉成 Struct

/*
//* target is the struct of model
func getJson(url string, target interface{}) error {
	//* myClient is http.Client
	resp, err := myClient.Get("")
	if err != nil {
			return err
	}

	//* remember to close the Body
	defer resp.Body.Close()

	//* create json Decoder and decode（可以縮寫成最下面一行）
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(target)
	if err != nil {
			return err
	}

	return nil

	//* json decode 的部分也可以縮寫成這一行
	//* return json.NewDecoder(r.Body).Decode(target)
}
*/


// ----------------------------------------------------------------------------
// *	Custom Date type with format YYYY-MM-DD and JSON decoder (Parser) and encoder (Unmarshal and Marshal methods)

//ISODate struct
type ISODate struct {
	Format string  `json:"format"`
  Time time.Time `json:"time"`
}

//UnmarshalJSON ISODate method
func (Date *ISODate) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	Date.Format = "2006-01-02"
	t, _ := time.Parse(Date.Format, s)
	Date.Time = t
	return nil
}

// MarshalJSON ISODate method
func (Date ISODate) MarshalJSON() ([]byte, error) {
	return json.Marshal(Date.Time.Format(Date.Format))
}

func TestCustomizedTimeFoermat(t *testing.T) {
	// *	struct to json
	date := ISODate{"2006-01-02", time.Now()}
	byteSlice, err := date.MarshalJSON()
	fmt.Println(string(byteSlice), err)
	// "2022-05-22" <nil>

	date2 := ISODate{"Mon Jan _2 15:04:05 2006", time.Now()}
	byteSlice2, err2 := date2.MarshalJSON()
	fmt.Println(string(byteSlice2), err2)
	// "Sun May 22 19:03:39 2022" <nil>


	// *	json to struct
	jsonString := `"2022-05-22"`
	// b := []byte{} // string to byte
	data4 := ISODate{}
	fmt.Println(data4) // { 0001-01-01 00:00:00 +0000 UTC}
	err = data4.UnmarshalJSON([]byte(jsonString))
	fmt.Println(err, data4) // <nil> {2006-01-02 0001-01-01 00:00:00 +0000 UTC}
	fmt.Printf("%+v\n", data4) // {Format:2006-01-02 Time:2022-05-22 00:00:00 +0000 UTC}
}

/*
	The function

	- time.Parse parses a date string, and
	- Format formats a time.Time.
	They have the following signatures:

	*		func Parse(layout, value string) (Time, error)
	*		func (t Time) Format(layout string) string
*/

func TestCustomizedTimeFormat(t *testing.T) {
	input := "2020-04-14"
	layout := "2006-01-02" // customized layout(time format), must same as input format
	ti, _ := time.Parse(layout, input)
	fmt.Println(ti) // 2020-04-14 00:00:00 +0000 UTC
	//* customized layout
	fmt.Println(ti.Format("02-Jan-2006")) // 14-Apr-2020
	//* golang自帶定義的layout
	fmt.Println(ti.Format(time.RFC850)) // Tuesday, 14-Apr-20 00:00:00 UTC

	t2, _ := time.Parse("2006-01-02", "2022-05-22")
	fmt.Println(t2)
}

// time.Parse 函式採用由時間格式佔位符作為第一個引數和代表時間的實際格式化字串作為第二個引數的佈局。
func TestTimeParse(t *testing.T) {
	//* customized layout
	timeT, _ := time.Parse("2006-01-02", "2020-04-14")
	fmt.Println(timeT)

	timeT, _ = time.Parse("06-01-02", "20-04-14")
	fmt.Println(timeT)

	timeT, _ = time.Parse("2006-Jan-02", "2020-Apr-14")
	fmt.Println(timeT)

	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05", "2020-Apr-14 Tuesday 03:19:25")
	fmt.Println(timeT)

	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05", "2020-Apr-14 Tuesday 23:19:25")
	fmt.Println(timeT)

	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05 PM MST -07:00", "2020-Apr-14 Tuesday 11:19:25 PM IST +05:30")
	fmt.Println(timeT)

	//* golang自帶定義的layout
	timeT, _ = time.Parse(time.Kitchen, "5:24AM")
	fmt.Println(timeT)

	timeT, _ = time.Parse(time.Stamp, "Nov 16 13:24:37")
	fmt.Println(timeT)

	// 2020-04-14 00:00:00 +0000 UTC
	// 2020-04-14 00:00:00 +0000 UTC
	// 2020-04-14 00:00:00 +0000 UTC
	// 2020-04-14 03:19:25 +0000 UTC
	// 0001-01-01 00:00:00 +0000 UTC	--> Wrong translation
	// 2020-04-14 23:19:25 +0530 IST
	// 0000-01-01 05:24:00 +0000 UTC
	// 0000-11-16 13:24:37 +0000 UTC
}