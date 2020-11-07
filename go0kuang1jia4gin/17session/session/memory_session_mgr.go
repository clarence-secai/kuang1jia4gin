package session

import (
	"crypto/rand"
	"fmt"
	"log"
	"sync"
)

//CreateUUID 生成UUID
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

func NewMemorySessionMgr() SessionMgr {
	sr := MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return &sr //写完下面的三个方法，这里就不会报错了，因为已经实现了接口要求的三个方法
}
func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	session = NewMemorySession(CreateUUID())
}
func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}
func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	return
}
