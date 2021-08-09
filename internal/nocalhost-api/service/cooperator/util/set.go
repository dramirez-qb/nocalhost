package util

import "sync"

type Set struct {
	inner sync.Map
}

func NewSet() *Set {
	return &Set{
		sync.Map{},
	}
}

func (s *Set) ToArray() []string {
	result := make([]string, 0)

	s.inner.Range(
		func(key, value interface{}) bool {
			result = append(result, key.(string))
			return true
		},
	)

	return result
}

func (s *Set) Put(key string) {
	s.inner.Store(key, "")
}

func (s *Set) Exist(key string) bool {
	_, ok := s.inner.Load(key)
	return ok
}