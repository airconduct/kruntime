package store

import "sync"

type Store[KT comparable, VT any] interface {
	Set(k KT, v VT)
	Get(k KT) (v VT, ok bool)
	Delete(k KT)
	Range(fn func(key KT, value VT) bool)
}

func New[KT comparable, VT any]() Store[KT, VT] {
	return &syncStore[KT, VT]{}
}

type syncStore[KT comparable, VT any] struct {
	m sync.Map
}

func (s *syncStore[KT, VT]) Set(k KT, v VT) {
	s.m.Store(k, v)
}

func (s *syncStore[KT, VT]) Get(k KT) (v VT, ok bool) {
	vv, ok := s.m.Load(k)
	if ok {
		return vv.(VT), ok
	}
	return v, ok
}

func (s *syncStore[KT, VT]) Delete(k KT) {
	s.m.Delete(k)
}

func (s *syncStore[KT, VT]) Range(fn func(key KT, value VT) bool) {
	s.m.Range(func(key, value any) bool {
		return fn(key.(KT), value.(VT))
	})
}
