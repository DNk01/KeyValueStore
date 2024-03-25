package store

import "time"

type KeyValueService interface {
	Set(key string, value interface{}, ttl time.Duration) error
	Get(key string) (value interface{}, err error)
	Delete(key string) error
}
