package oncesingleflight

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCollection(t *testing.T) {
	m := require.New(t)

	var col Collection

	key := 1

	elem1 := col.Get(key)
	elem2 := col.Get(key)

	m.Equal(elem1, elem2)
	m.False(elem1.Finished)

	shared1 := elem1.Do(func() interface{} {
		return 2
	})
	m.False(shared1)

	shared2 := elem2.Do(func() interface{} {
		return 3
	})
	m.True(shared2)

	m.True(elem1.Finished)
	m.Equal(2, elem1.Result)
	m.Equal(2, elem2.Result)
}

func TestCollection_Panic(t *testing.T) {
	m := require.New(t)

	var col Collection

	key := 1

	elem1 := col.Get(key)
	elem2 := col.Get(key)

	func() {
		defer func() {
			m.Equal(2, recover())
		}()
		shared1 := elem1.Do(func() interface{} {
			panic(2)
		})
		m.False(shared1)
	}()

	func() {
		defer func() {
			m.Nil(recover())
		}()
		shared2 := elem2.Do(func() interface{} {
			panic(3)
		})
		m.True(shared2)
	}()

	m.False(elem1.Finished)
	m.Nil(elem1.Result)
	m.Nil(elem2.Result)
}
