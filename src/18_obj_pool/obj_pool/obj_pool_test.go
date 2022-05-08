package objpool

import (
	"fmt"
	"testing"
	"time"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	// Overflow --> due to put improper type data in pool
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// 	t.Error(err)
	// }
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			// if don't release request 11 but only 10 in objpool --> the last one request timeout
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Printf("Done")

	// *objpool.ReusableObj
	// *objpool.ReusableObj
	// *objpool.ReusableObj
	// *objpool.ReusableObj
	// *objpool.ReusableObj
	// *objpool.ReusableObj
	// *objpool.ReusableObj
	// *objpool.ReusableObj
	// *objpool.ReusableObj
	// *objpool.ReusableObj
	// *objpool.ReusableObj
}