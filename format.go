package dlt645

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 压缩BCD码解析成float64 bcd-小字序BCD字节 format-格式"XXX.XXX" allowNegative-是否需要处理负号
func BcdBytes2Float64(bcd []byte, format string, allowNegative bool) (float64, error) {
	// 判断format中"X"数量与bcd是否一致
	xCount := strings.Count(format, "X")
	if xCount != len(bcd)*2 {
		return 0, errors.New("format length is not equal to bcd length")
	}

	// 将BCD中的小字序字节转为大字序字节
	bcd = BytesReverse(bcd)
	// 判断是否为负数
	var isNegative bool
	if allowNegative {
		isNegative = bcd[0]&0x80 > 0
		bcd[0] = bcd[0] & 0x7F
	}

	// bcd转换成字符串
	bcdString := ""
	for _, value := range bcd {
		bcdString += fmt.Sprintf("%02X", value)
	}

	// 将BCD的数值赋值到format中
	resultString := ""
	bcdIndex := 0
	for _, value := range format {
		if value == 'X' {
			resultString += string(bcdString[bcdIndex])
			bcdIndex++
		} else {
			resultString += string(value)
		}
	}

	// 将字符串转换为float64
	result, err := strconv.ParseFloat(resultString, 64)
	if err != nil {
		return 0, err
	}

	if isNegative {
		result = -result
	}
	return result, nil
}
