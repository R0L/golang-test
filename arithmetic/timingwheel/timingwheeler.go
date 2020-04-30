package timingwheel

import "container/list"

// 时间轮标准
type TimingWheeler interface {
	AddTask(delay int, task Task) (string, error)
	RemoveTask(taskNo string) error
	Execute(tasks *TaskList) error
	Promote() error
}

// 任务
type Task func() error

// 任务链表
type TaskList struct {
	index int
	list  *list.List
}
