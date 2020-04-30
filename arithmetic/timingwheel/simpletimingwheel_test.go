package timingwheel

import (
	"testing"
)

// 测试简单时间轮推动
func TestSimpleTimingWheel_Promote(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			"测试简单时间轮推动",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSimpleTimingWheel(60)
			if err := s.Promote(); (err != nil) != tt.wantErr {
				t.Errorf("Promote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
