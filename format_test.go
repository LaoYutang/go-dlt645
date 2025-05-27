package dlt645_test

import (
	"testing"

	"github.com/LaoYutang/go-dlt645"
)

func TestBcdBytes2Float64(t *testing.T) {
	tests := []struct {
		name          string
		bcd           []byte
		digit         uint8
		allowNegative bool
		expected      float64
	}{
		{"Valid conversion 1", []byte{0x12, 0x94}, 2, false, 94.12},
		{"Valid conversion 2", []byte{0x12, 0x94}, 2, true, -14.12},
		{"Valid conversion 3", []byte{0x12, 0x34, 0x56}, 4, false, 56.3412},
		{"Valid conversion 4", []byte{0x12, 0x34, 0x56, 0x78}, 2, false, 785634.12},
		{"Valid conversion 5", []byte{0x5, 0x96}, 2, false, 96.05},
		{"Valid conversion 5", []byte{0x5, 0x96}, 2, true, -16.05},
		{"Valid conversion 6", []byte{0x34, 0x1, 0xA1}, 1, true, -21013.4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := dlt645.BcdBytes2Float64(tt.bcd, tt.digit, tt.allowNegative)
			// 正常测试但未返回正确结果
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
