package dlt645

import "strconv"

// 字符串转BCD字节
func String2BcdBytes(str string) []byte {
	if len(str)&1 == 1 {
		str = "0" + str
	}
	res := make([]byte, 0, len(str)/2)
	for i := 0; i < len(str); i += 2 {
		high := min(str[i]-'0', 9)
		low := min(str[i+1]-'0', 9)
		res = append(res, high<<4|low)
	}
	return res
}

func String2BcdBytes2(str string) []byte {
	slen := len(str)
	bHex := make([]byte, len(str)/2)
	ii := 0
	for i := 0; i < len(str); i = i + 2 {
		if slen != 1 {
			ss := string(str[i]) + string(str[i+1])
			bt, _ := strconv.ParseInt(ss, 16, 32)
			bHex[ii] = byte(bt)
			ii = ii + 1
			slen = slen - 2
		}
	}
	return bHex
}

// BCD字节转字符串
func BcdBytes2String(bs []byte) string {
	if len(bs) == 0 {
		return ""
	}

	result := make([]byte, len(bs)*2)
	const hexChars = "0123456789abcdef"

	for i, b := range bs {
		result[i*2] = hexChars[b>>4]
		result[i*2+1] = hexChars[b&0x0f]
	}
	return string(result)
}

// 字节切片全部减小
func BytesSub(bs []byte, sub byte) []byte {
	r := []byte{}
	for _, b := range bs {
		r = append(r, b-sub)
	}
	return r
}

// 字节切片全部增加
func BytesAdd(bs []byte, add byte) []byte {
	r := []byte{}
	for _, b := range bs {
		r = append(r, b+add)
	}
	return r
}

// 反转字节切片
func BytesReverse(bs []byte) []byte {
	r := make([]byte, len(bs))
	for i, b := range bs {
		r[len(bs)-i-1] = b
	}
	return r
}
