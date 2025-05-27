package dlt645

import (
	"math"
)

// 压缩BCD码解析成float64 bcd-小字序BCD字节 format-格式"XXX.XXX" allowNegative-是否需要处理负号
func BcdBytes2Float64(data []byte, digit uint8, allowNegative bool) (result float64) {
	for i, value := range data {
		if allowNegative && i == len(data)-1 {
			if value&0x80 != 0 {
				value = value & 0x7F
			} else {
				allowNegative = false
			}
		}
		result += (float64(value>>4)*10 + float64(value&0x0F)) * math.Pow(10, float64(i*2))
	}

	result /= math.Pow(10, float64(digit))

	if allowNegative {
		result = -result
	}

	return
}
