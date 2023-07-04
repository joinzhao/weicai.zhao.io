package syncx

import (
	"math"
	"sync"
	"time"
)

// LeakyBucket 漏桶
type LeakyBucket struct {
	rate        float64    // 速率 每秒出水量， 秒为最小计时单位
	cap         float64    // 容量
	water       float64    // 当前桶内容量
	lastLeakyMs int64      // 上次漏水时间
	mu          sync.Mutex // 锁
}

func NewLeakyBucket(cap, rate float64) *LeakyBucket {
	return &LeakyBucket{cap: cap, rate: rate}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	// 计算出水量
	now := time.Now().UnixNano()                                            // 当前时间
	eclipse := float64(now-lb.lastLeakyMs) / float64(time.Second) * lb.rate // 上次漏水-当前时间应该漏水量
	lb.water = math.Max(lb.water-eclipse, 0)                                // 当前水量
	lb.lastLeakyMs = now

	if (lb.water + 1) < lb.cap {
		lb.water++
		return true
	}
	return false
}

func (lb *LeakyBucket) Set(cap, rate float64) {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	lb.cap = cap
	lb.rate = rate
}
