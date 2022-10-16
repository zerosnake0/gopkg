package oncesingleflight

import "sync"

type Callback func() interface{}

type Element struct {
	parent   *Collection
	key      interface{}
	once     sync.Once
	Result   interface{}
	Finished bool // not panicked
}

func (e *Element) Do(callback Callback) bool {
	shared := true
	e.once.Do(func() {
		defer e.parent.m.Delete(e.key)
		e.Result = callback()
		shared = false
		e.Finished = true
	})
	return shared
}
