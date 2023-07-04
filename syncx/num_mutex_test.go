package syncx_test

import (
	"fmt"
	"sync"
	"testing"
	"weicai.zhao.io/syncx"
)

func TestMutex(t *testing.T) {
	var mu = syncx.NumMutex{Max: 5}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		var j = i
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()
			//time.Sleep(time.Second)
			fmt.Println("now: ", i)
			wg.Done()
		}(j)
	}

	wg.Wait()
}
