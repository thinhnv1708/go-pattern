package lib

import "sync"

var (
	instance *singleton
	once     sync.Once
)

type Singleton interface {
	AddOne() int
}

type singleton struct {
	counter int
}

func (s *singleton) AddOne() int {
	s.counter += 1

	return s.counter
}

func GetInstance() *singleton {

	once.Do(func() {
		instance = &singleton{counter: 100}
	})

	return instance
}
