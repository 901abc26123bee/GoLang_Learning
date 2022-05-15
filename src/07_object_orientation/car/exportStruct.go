package car

// 匯出的欄位（Exported fields）
// 如同 package 中的變數一樣，struct 中的欄位只有在 `欄位名稱` 以大寫命名時才會 export 出去，其他 package 中才能取用得到：

type Car struct {
	Name  string
	price float32 // cannot export price --> need to change to Price
	amount int8
}

type Toyota struct {
	Name  string
	Price float32
	Amount int8
}