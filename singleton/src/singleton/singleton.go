package singleton

import (
	"encoding/json"
	"sync"
)

var (
	one singleton
)

func GetSingleton() *singleton {
	one.init()
	return &one
}

type singleton struct {
	Name string

	once sync.Once
}

func (s *singleton) init() {
	s.once.Do(
		func() {
			s.Name = "one"
		})
}

func (s *singleton) String() string {
	bytes, _ := json.Marshal(s)
	return string(bytes)
}
