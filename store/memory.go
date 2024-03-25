package store

import (
	"container/heap"
	"log"
	"sync"
	"time"
)

type KeyValueStore struct {
	items map[string]*Item
	mu    sync.Mutex
	pq    PriorityQueue
}
type Item struct {
	Key       string
	Value     interface{}
	ExpiresAt time.Time
}

func NewKVStore() *KeyValueStore {
	kv := &KeyValueStore{
		items: make(map[string]*Item),
		pq:    make(PriorityQueue, 0),
	}
	heap.Init(&kv.pq)
	go kv.cleanupExpiredItems()
	return kv
}

func (kv *KeyValueStore) Set(key string, value interface{}, ttl time.Duration) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	expiresAt := time.Now().Add(ttl)
	item := &Item{
		Key:       key,
		Value:     value,
		ExpiresAt: expiresAt,
	}

	log.Println("Setting key", key, "with TTL", ttl)
	kv.items[key] = item
	heap.Push(&kv.pq, item)
	return nil
}

func (kv *KeyValueStore) Get(key string) (interface{}, error) {
	if kv.items[key] == nil {
		log.Println("Key not found", key)
		return nil, nil
	}
	log.Println("Getting key", key)
	return kv.items[key].Value, nil
}

func (kv *KeyValueStore) Delete(key string) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	log.Println("Deleting key", key)
	delete(kv.items, key)
	return nil
}

func (kv *KeyValueStore) cleanupExpiredItems() {
	ticker := time.NewTicker(3 * time.Second)
	for {
		select {
		case <-ticker.C:
			kv.removeExpired()
		}
	}
}

func (kv *KeyValueStore) removeExpired() {
	kv.mu.Lock()
	defer kv.mu.Unlock()

	now := time.Now()
	for kv.pq.Len() > 0 {
		item := kv.pq[0]
		if item.ExpiresAt.After(now) {
			break
		}
		log.Println("Removing expired item with key", item.Key)
		heap.Pop(&kv.pq)
		delete(kv.items, item.Key)
	}
}
