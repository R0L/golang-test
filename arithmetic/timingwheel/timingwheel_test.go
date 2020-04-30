package timingwheel

import "testing"

// 测试时间轮开始
func TestTimingWheel_start(t *testing.T) {
	type fields struct {
		TimingWheeler TimingWheeler
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			"测试时间轮开始",
			fields{TimingWheeler: NewSimpleTimingWheel(60)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TimingWheel{
				TimingWheeler: tt.fields.TimingWheeler,
			}
			s.start()
		})
	}
}
