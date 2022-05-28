package context_test

import ()

/*

	* context package 最重要的就是處理多個 goroutine 的情況，特別是用來送出取消或結束的 signal。

		我們可以在一個 goroutine 中建立 Context 物件後，傳入另一個 goroutine；
		另一個 goroutine 即可以透過 Done() 來從該 context 中取得 signal，
		一旦這個 Done channel 關閉之後，這個 goroutine 即會關閉並 return。

		Context 也可以是受時間控制，它也可以在特定時間後關閉該 signal channel，
		我們可以定義一個 deadline 或 timeout 的時間，時間到了之後，Context 物件就會關閉該 signal channel。

		更好的是，一旦父層的 Context 關閉其 Done channel 之後，子層的 Done channel 則會自動關閉。

*/

/*
	重要概念
	* 不要把 Context 保存在 struct 中，而是直接當作第一個參數傳入 function 或 goroutine 中，通常會命名為 ctx
	* server 在處理傳進來的請求時應該要建立一個 Context，而使用該 server 的方法則應該要接收 Context 作為參數
	* 雖然函式可以允許傳入 nil Context，但千萬不要這麼做，如果你不確定要用哪個 Context，可以使用 context.TODO
	* 只在 request-scoped data 這種要交換處理資料或 API 的範疇下使用 context Values，不要傳入 optional parameters 到函式中。
	* 相同的 Context 可以傳入多個不同的 goroutine 中使用，在多個 goroutines 中同時使用 Context 是安全的（safe）
*/