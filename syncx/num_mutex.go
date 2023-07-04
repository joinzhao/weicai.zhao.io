package syncx

import "sync"

// NumMutex 计数锁
type NumMutex struct {
	Max     uint32
	limit   uint32
	mu      sync.Mutex
	lockTag bool
}

func (m *NumMutex) Lock() {
	m.mu.Lock()
	m.lockTag = true

	now := m.limit + 1
	if now <= m.Max {
		m.limit = now
		m.lockTag = false
		m.mu.Unlock()
		return
	}
}
func (m *NumMutex) Unlock() {
	// 限制达到上限，已经上锁，只执行解锁， 其他情况不操作
	if m.lockTag {
		m.limit -= 1
		m.lockTag = false
		m.mu.Unlock()
		return
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	m.limit -= 1
	if m.limit < 0 {
		m.limit = 0
	}
}
