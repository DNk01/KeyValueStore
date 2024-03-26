package test

import (
	"github.com/stretchr/testify/assert"
	"keyValueStorage/store"
	"testing"
	"time"
)

func TestMemoryStore_SetAndGet(t *testing.T) {
	memoryStore := store.NewKVStore()
	key := "testKey"
	value := "testValue"
	ttl := time.Second * 10

	err := memoryStore.Set(key, value, ttl)
	assert.NoError(t, err)

	gotValue, err := memoryStore.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, value, gotValue)
}

func TestMemoryStore_ExpiredKey(t *testing.T) {
	memoryStore := store.NewKVStore()
	key := "expiredKey"
	value := "value"
	ttl := time.Millisecond * 100

	err := memoryStore.Set(key, value, ttl)
	assert.NoError(t, err)

	time.Sleep(ttl + time.Millisecond*3000) //TODO Need to change when ticker time is changed

	result, _ := memoryStore.Get(key)
	assert.Equal(t, nil, result)
}

func TestMemoryStore_Delete(t *testing.T) {
	memoryStore := store.NewKVStore()
	key := "deleteKey"
	value := "value"

	err := memoryStore.Set(key, value, 0)
	assert.NoError(t, err)
	err = memoryStore.Delete(key)
	assert.NoError(t, err)

	result, _ := memoryStore.Get(key)
	assert.Equal(t, nil, result)
}
