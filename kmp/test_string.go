package kmp

var test_str = `package rwlock

import (
	"sync"
)

type RWLocker interface {
	RLock()
	RUnlock()
	Lock()
	Unlock()
}

type ReadCountRWLocker struct {
	rm *sync.Mutex
	wm *sync.Mutex
	c  int
}

func NewReadCountRWLocker() *ReadCountRWLocker {
	locker := &ReadCountRWLocker{
		rm: &sync.Mutex{},
		wm: &sync.Mutex{},
	}
	return locker
}

func (rw *ReadCountRWLocker) Lock() {
	rw.wm.Lock()
}

func (rw *ReadCountRWLocker) Unlock() {
	rw.wm.Unlock()
}

func (rw *ReadCountRWLocker) RLock() {
	defer rw.rm.Unlock()
	rw.rm.Lock()
	// 确认没有读，lock写锁
	if rw.c == 0 {
		rw.c += 1
		rw.wm.Lock()
	}
}

func (rw *ReadCountRWLocker) RUnlock() {
	defer rw.rm.Unlock()
	rw.rm.Lock()
	// 确认是最后一个读，unlock写锁
	if rw.c == 1 {
		rw.c -= 1
		rw.wm.Unlock()
	}
}

type RWCondLocker struct {
	m              *sync.Mutex
	c              int // -1表示有写操作，0表示无操作，大于0表示有读操作
	readCond       *sync.Cond
	writeCond      *sync.Cond
	readWaitCount  int
	writeWaitCount int
}

func NewRWCondLocker() *RWCondLocker {
	locker := &RWCondLocker{
		m: &sync.Mutex{},
	}
	locker.readCond = sync.NewCond(locker.m)
	locker.writeCond = sync.NewCond(locker.m)
	return locker
}

func (rw *RWCondLocker) RLock() {
	rw.m.Lock()

	for rw.c < 0 {
		rw.readWaitCount += 1
		rw.readCond.Wait()
		rw.readWaitCount -= 1
	}

	rw.c += 1

	rw.m.Unlock()
}

func (rw *RWCondLocker) WLock() {
	rw.m.Lock()

	for rw.c > 0 || rw.c == -1 {
		rw.writeWaitCount += 1
		rw.writeCond.Wait()
		rw.writeWaitCount -= 1
	}

	rw.c -= 1

	rw.m.Unlock()
}

func (rw *RWCondLocker) Unlock() {
	rw.m.Lock()
	defer rw.m.Unlock()

	if rw.c == -1 {
		rw.c += 1
	} else if rw.c > 0 {
		rw.c -= 1
	}

	if rw.c > 0 {
		if rw.readWaitCount > 0 {
			rw.readCond.Signal()
		}
	} else if rw.c == 0 {
		if rw.writeWaitCount > 0 {
			rw.writeCond.Signal()
		}
		if rw.readWaitCount > 0 {
			rw.readCond.Broadcast()
		}
	}
}

func (rw *RWCondLocker) TryRLock() bool {
	ret := false
	rw.m.Lock()
	// 没有写操作时
	if rw.c >= 0 {
		ret = true
	}
	rw.c += 1
	rw.m.Unlock()
	return ret
}

func (rw *RWCondLocker) TryWLock() bool {
	ret := false
	rw.m.Lock()
	// 没有任何读写操作时
	if rw.c == 0 {
		ret = true
	}
	rw.c -= 1
	rw.m.Unlock()
	return ret
}`
