package aes

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

// 测试加密
func TestEncrypt(t *testing.T) {
	type args struct {
		src string
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"测试加密",
			args{
				src: "1591016258353_2_R022005310000520066",
				key: "1234567890123456",
			},
			"VmiEoAnpc+4yk1/ntWKNTFsPt70MtqhpIxZ1f/fVoAd5MrTN+90Q8/fpygVtGJQB",
			false,
		},
		{
			"测试加密",
			args{
				src: fmt.Sprintf("%d_2_R022005310000520100", time.Now().UnixNano()),
				key: "1234567890123456",
			},
			"VmiEoAnpc+4yk1/ntWKNTFsPt70MtqhpIxZ1f/fVoAd5MrTN+90Q8/fpygVtGJQB",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAES(tt.args.key).Encrypt(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(string(got[:]))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// 测试解密
func TestDecrypt(t *testing.T) {
	type args struct {
		src string
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			"测试解密",
			args{
				src: "VmiEoAnpc+4yk1/ntWKNTFsPt70MtqhpIxZ1f/fVoAd5MrTN+90Q8/fpygVtGJQB",
				key: "1234567890123456",
			},
			"1591016258353_2_R022005310000520066",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAES(tt.args.key).Decrypt(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() got = %v, want %v", got, tt.want)
			}
		})
	}
}
