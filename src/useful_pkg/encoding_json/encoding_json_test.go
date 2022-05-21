package encodingjson_test

import ()

// json.Unmarshal 把 json 字串轉成 struct
// json.Marshal 把 struct 轉成 json 字串

// -------------------------- 將 struct 轉成 JSON --------------------------
// 定義可轉換成 JSON 的 struct
// 透過 struct field's tag 可以將 struct 轉換成特定的格式：

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
omitempty ：如果該 field 是空值，則忽略該 field 的欄位，空值包含 false, 0, nil pointer, nil interface value, 空陣列、空字串
string：如果某欄位的值又是 JSON 的話，可以在 options 中使用 string
*/