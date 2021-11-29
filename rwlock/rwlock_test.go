package rwlock

import (
	"context"
	"testing"
	"time"

	"github.com/huoshan017/gopool"
)

/*func TestReadCountRWLocker(t *testing.T) {
	rw_locker := NewReadCountRWLocker()
	count := 10000
	go func() {
		for i := 0; i < 100; i++ {
			go func() {
				rw_locker.Lock()
				count -= 1
				rw_locker.Unlock()
			}()
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			go func() {
				rw_locker.RLock()
				t.Logf("count = %v", count)
				rw_locker.RUnlock()
			}()
		}
	}()

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
	}
}*/

func TestRWCountLocker(t *testing.T) {
	pool := gopool.NewPool(1000)
	rw_locker := NewRWCondLocker()
	count := 10000

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		for i := 0; i < 300000; i++ {
			pool.CommitTask(ctx, func(_ interface{}) {
				rw_locker.WLock()
				count -= 1
				t.Logf("write count from %v to %v", count+1, count)
				rw_locker.Unlock()
			}, nil)
		}
	}()

	go func() {
		for i := 0; i < 300000; i++ {
			pool.CommitTask(ctx, func(_ interface{}) {
				rw_locker.WLock()
				count += 1
				t.Logf("write count from %v to %v", count-1, count)
				rw_locker.Unlock()
			}, nil)
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			pool.CommitTask(ctx, func(_ interface{}) {
				rw_locker.RLock()
				t.Logf("read count = %v", count)
				rw_locker.Unlock()
			}, nil)
		}
	}()

	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
	}
}
