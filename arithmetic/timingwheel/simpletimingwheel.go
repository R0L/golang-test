package timingwheel

import (
	"container/list"
	"container/ring"
	"fmt"
	"github.com/satori/go.uuid"
	"time"
)

type SimpleTimingWheel struct {
	slot  *ring.Ring
	index int
}

type taskNode struct {
	taskNo string
	delay  int
	task   Task
	err    error
}

// 初始化
func NewSimpleTimingWheel(slot int) TimingWheeler {
	return &SimpleTimingWheel{
		slot: ring.New(slot),
	}
}

func (s *SimpleTimingWheel) LocationTask(index int, delay int, task Task) (string, error) {
	move := index % s.slot.Len()
	var tasks *list.List
	slot := s.slot.Move(move)
	if slot.Value == nil {
		tasks = list.New()
	} else {
		tasks = slot.Value.(*list.List)
	}
	taskNo := uuid.Must(uuid.NewV4(), nil)
	tasks.PushBack(&taskNode{
		taskNo: taskNo.String(),
		delay:  delay,
		task:   task,
	})
	slot.Value = tasks

	return taskNo.String(), nil
}

func (s *SimpleTimingWheel) AddTask(delay int, task Task) (string, error) {
	return s.LocationTask(s.index+delay, delay, task)
}

func (s *SimpleTimingWheel) RemoveTask(taskNo string) error {
	panic("implement me")
}

// 执行任务
func (s *SimpleTimingWheel) Execute(tasks *TaskList) error {
	for {
		element := tasks.list.Front()
		// 链表为空
		if nil == element {
			return nil
		}
		node, ok := element.Value.(*taskNode)
		if !ok {
			return nil
		}
		// 计算此任务的下一个周期
		_, err := s.LocationTask(node.delay+tasks.index-s.index, node.delay, node.task)
		if nil != err {
			return err
		}
		// 删除此任务
		tasks.list.Remove(element)
		go func() {
			// todo 触发任务系统
			fmt.Println(node.taskNo)
		}()
	}
}

// 推动时间轮
func (s *SimpleTimingWheel) Promote() error {
	ticker := time.NewTicker(1 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			s.slot = s.slot.Next()
			s.index++
			s.index %= s.slot.Len()
			go func(index int) {
				if tasks, ok := s.slot.Value.(*list.List); ok {
					task := s.Execute(&TaskList{index: index, list: tasks})
					if nil != task {
						fmt.Println(task)
					}
				}
			}(s.index)
			fmt.Println(".", s.index)
		}
	}
}
