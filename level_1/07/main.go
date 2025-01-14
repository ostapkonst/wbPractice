// 7. Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

type SyncStrMap struct {
	value map[string]string
	mutex *sync.RWMutex
}

func NewSyncStrMap() *SyncStrMap {
	return &SyncStrMap{
		value: make(map[string]string),
		mutex: &sync.RWMutex{},
	}
}

func (sm *SyncStrMap) Get(key string) string {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	return sm.value[key]
}

func (sm *SyncStrMap) Set(key string, value string) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sm.value[key] = value
}

func main() {
	sm := NewSyncStrMap()
	sm.Set("key", "value")
	fmt.Println("result:", sm.Get("key"))
}
