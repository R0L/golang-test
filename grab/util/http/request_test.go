package http

import "testing"

// 测试请求
func Test_request(t *testing.T) {
	type args struct {
		url    string
		rParam *RequestParam
		result interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"测试请求Map",
			args{
				url: "https://www.baidu.com",
				rParam: &RequestParam{
					Method:      "get",
					ContentType: "application/x-www-form-urlencoded",
					Cookie:      "",
					Param: &map[string]interface{}{
						"a": "A",
						"b": "B",
						"c": "C",
					},
				},
				result: nil,
			},
		},
		{
			"测试请求Struct",
			args{
				url: "https://www.baidu.com",
				rParam: &RequestParam{
					Method:      "get",
					ContentType: "application/x-www-form-urlencoded",
					Cookie:      "",
					Param: struct {
						A int
						B string
					}{
						A: 10,
						B: "bb",
					},
				},
				result: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewHttp().Request(tt.args.url, tt.args.rParam, tt.args.result)
		})
	}
}
