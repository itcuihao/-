package main

import "sync"

type Lock struct {
	c chan struct{}
}

func NewLock() Lock {
	var l Lock
	// 利用size为1的chan做lock
	l.c = make(chan struct{}, 1)
	l.c <- struct{}{}
	return l
}

func (l Lock) Lock() bool {
	lockResult := false
	select {
	case <-l.c:
		lockResult = true
	default:
	}
	return lockResult
}

func (l Lock) Unlock() {
	l.c <- struct{}{}
}

var counter int

func main() {
	var l = NewLock()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !l.Lock() {
				println("lock failed")
				return
			}
			counter++
			println("current counter ", counter)
			l.Unlock()
		}()
	}
	wg.Wait()
}
