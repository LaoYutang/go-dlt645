package dlt645_test

import (
	"testing"

	"github.com/LaoYutang/go-dlt645"
)

func TestBcdBytes2Float64(t *testing.T) {
	tests := []struct {
		name          string
		bcd           []byte
		format        string
		allowNegative bool
		expected      float64
		expectError   bool
	}{
		{
			name:          "Valid conversion 1",
			bcd:           []byte{0x12, 0x94},
			format:        "XX.XX",
			allowNegative: false,
			expected:      94.12,
			expectError:   false,
		},
		{
			name:          "Valid conversion 2",
			bcd:           []byte{0x12, 0x94},
			format:        "XX.XX",
			allowNegative: true,
			expected:      -14.12,
			expectError:   false,
		},
		{
			name:          "Valid conversion 3",
			bcd:           []byte{0x12, 0x34, 0x56},
			format:        "XX.XXXX",
			allowNegative: false,
			expected:      56.3412,
			expectError:   false,
		},
		{
			name:          "Valid conversion 4",
			bcd:           []byte{0x12, 0x34, 0x56, 0x78},
			format:        "XXXXXX.XX",
			allowNegative: false,
			expected:      785634.12,
			expectError:   false,
		},
		{
			name:          "Valid conversion 5",
			bcd:           []byte{0x5, 0x96},
			format:        "XX.XX",
			allowNegative: true,
			expected:      -16.05,
			expectError:   false,
		},
		{
			name:          "Valid conversion 6",
			bcd:           []byte{0x34, 0x1, 0xA1},
			format:        "XXXX.XX",
			allowNegative: true,
			expected:      -2101.34,
			expectError:   false,
		},
		{
			name:          "X count less than required",
			bcd:           []byte{0x12},
			format:        "XXX",
			allowNegative: false,
			expected:      0,
			expectError:   true,
		},
		{
			name:          "X count more than required",
			bcd:           []byte{0x12, 0x34},
			format:        "X",
			allowNegative: false,
			expected:      0,
			expectError:   true,
		},
		{
			name:          "Invalid float string",
			bcd:           []byte{0x1A},
			format:        "XX",
			allowNegative: false,
			expected:      0,
			expectError:   true,
		},
		{
			name:          "Single byte BCD",
			bcd:           []byte{0x01},
			format:        "X.X",
			allowNegative: false,
			expected:      0.1,
			expectError:   false,
		},
		{
			name:          "Byte order reversal",
			bcd:           []byte{0xCD, 0xAB},
			format:        "XXXX",
			allowNegative: false,
			expected:      0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := dlt645.BcdBytes2Float64(tt.bcd, tt.format, tt.allowNegative)
			// 错误测试但返回正常结果
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}
				return
			}
			// 正常测试但返回错误
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}
			// 正常测试但未返回正确结果
			if !tt.expectError && result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
