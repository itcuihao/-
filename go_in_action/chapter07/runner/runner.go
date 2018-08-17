package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

// Runner 在给定的时间内执行一组任务
// 并且在操作系统发送中断信号时结束这些任务
type Runner struct {
	interrupt chan os.Signal

	complete chan error

	timeout <-chan time.Time

	tasks []func(int)
}

// ErrTimeout 会在任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// ErrInterrupt 在接收操作系统事件时返回
var ErrInterrupt = errors.New("received interrupt")

// NewR 新建一个Runner
func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

// Add 添加任务
func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start 执行所有任务，并监视通道事件
func (r *Runner) Start() error {

	// 接收所有中断信号
	signal.Notify(r.interrupt, os.Interrupt)

	// 用不同的goroutine执行不同的任务
	go func() {
		r.complete <- r.run()
	}()

	select {
	// 任务处理完成
	case err := <-r.complete:
		return err

	// 发出超时信号
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {

		// 检测中断信号
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行注册任务
		task(id)
	}

	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	// 正常
	default:
		return false
	// 中断事件被触发
	case <-r.interrupt:
		// 停止接受后续任何信号
		signal.Stop(r.interrupt)
		return true
	}
}
