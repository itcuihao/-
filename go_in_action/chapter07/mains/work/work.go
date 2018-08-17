package main

import (
	"goalgorithm/go_in_action/chapter07/work"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"therese",
	"jason",
}

// 使用特定方式打印名字
type namePrinter struct {
	name string
}

// Task实现Worker接口
func (m *namePrinter) Task() {
	log.Println(m.name)
	time.Sleep(time.Second)
}

func main() {
	p := work.New(2)

	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		// 迭代names切片
		for _, n := range names {
			np := namePrinter{
				name: n,
			}
			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()

	p.Shutdown()
}
