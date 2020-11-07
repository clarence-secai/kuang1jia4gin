package session

//实现了session.go中的Session接口
import (
	"errors"
	"sync"
)

type MemorySession struct {
	sessionId string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

func NewMemorySession(id string) *MemorySession {
	s := MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
	return &s
}
func (m *MemorySession) Set(key string, value interface{}) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key] = value
	return
}
func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("该session的key不存在")
		return
	}
	return
}
func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key) //delete是官方文档中的内建函数
	return
}
func (m *MemorySession) Save() (err error) {
	return
}
