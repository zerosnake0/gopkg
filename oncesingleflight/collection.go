package oncesingleflight

import (
	"sync"
)

type Collection struct {
	m sync.Map
}

func (col *Collection) Get(key interface{}) *Element {
	newV := &Element{
		parent: col,
		key:    key,
	}
	actual, _ := col.m.LoadOrStore(key, newV)
	return actual.(*Element)
}

func (col *Collection) Do(key interface{}, callback Callback) (*Element, bool) {
	e := col.Get(key)
	shared := e.Do(callback)
	return e, shared
}
