package lrn

import (
	"container/list"
	"fmt"
	"testing"
)

// 测试lrn
func TestLRN_Lrn(t *testing.T) {
	type fields struct {
		len  int
		list *list.List
	}
	type args struct {
		keys []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"测试lrn",
			fields{
				len: 10,
			},
			args{
				keys: []string{
					"1",
					"2",
					"3",
					"4",
					"5",
					"6",
					"7",
					"8",
					"9",
					"10",
					"11",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewLRN(tt.fields.len)
			if err := s.Lrn(tt.args.keys...); (err != nil) != tt.wantErr {
				t.Errorf("Lrn() error = %v, wantErr %v", err, tt.wantErr)
			}
			dump(s.list)
		})
	}
}

func dump(list *list.List) {
	for ele := list.Front(); ele != nil; {
		fmt.Println(ele.Value)
		ele = ele.Next()
	}
}
