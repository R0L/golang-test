package timingwheel

import "fmt"

type TimingWheel struct {
	TimingWheeler
}

func (s *TimingWheel) start() {

	s.TimingWheeler.AddTask(5, func() error {
		fmt.Println("5 delay")
		return nil
	})

	s.TimingWheeler.AddTask(15, func() error {
		fmt.Println("15 delay")
		return nil
	})

	// 启动时间轮
	if err := s.TimingWheeler.Promote(); nil != err {
		fmt.Println(err)
	}

}
