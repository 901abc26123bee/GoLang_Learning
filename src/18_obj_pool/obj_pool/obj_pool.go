package objpool

import (
	"errors"
	"time"
)

type ReusableObj struct {

}

type ObjPool struct {
	bufChan chan *ReusableObj
}

type ObjPoolWithAnyType struct {
	bufChan chan interface{}
}

func NewObjPool(numOgObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ReusableObj, numOgObj)
	for i := 0; i < numOgObj; i++ {
		objPool.bufChan <- &ReusableObj{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

// if you want to accept any type in to buffer channel --> obj interface{}
// but you need to do type determination when comsumed data from channel
func (p *ObjPool) GetObjWithAnyType(timeout time.Duration) (obj interface{}, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
		case p.bufChan <- obj:
			return nil
		default:
			return errors.New("Overflow") // put unfit type into channel --> blocked
	}
}