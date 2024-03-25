package mock

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type KeyValueService struct {
	mock.Mock
}

func (m *KeyValueService) Set(key string, value interface{}, ttl time.Duration) error {
	args := m.Called(key, value, ttl)
	return args.Error(0)
}

func (m *KeyValueService) Get(key string) (interface{}, error) {
	args := m.Called(key)
	return args.Get(0), args.Error(1)
}

func (m *KeyValueService) Delete(key string) error {
	args := m.Called(key)
	return args.Error(0)
}
