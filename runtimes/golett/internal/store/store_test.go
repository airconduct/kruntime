package store

import (
	"fmt"
	"testing"
)

type Foo struct {
	A string
}

func TestStore(t *testing.T) {
	s := New[string, *Foo]()
	s.Set("foo1", &Foo{A: "1"})
	s.Set("foo2", &Foo{A: "2"})
	s.Range(func(key string, value *Foo) bool {
		fmt.Println(value.A)
		return true
	})
}
