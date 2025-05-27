package dlt645

// 字符串转BCD字节
func String2BcdBytes(str string) []byte {
	if len(str)&1 == 1 {
		str = "0" + str
	}
	res := make([]byte, len(str)/2)
	for i := 0; i < len(str); i += 2 {
		high := charToHex(str[i])
		low := charToHex(str[i+1])
		res[i/2] = high<<4 | low
	}
	return res
}

// 字符转十六进制值
func charToHex(c byte) byte {
	switch {
	case c >= '0' && c <= '9':
		return c - '0'
	case c >= 'A' && c <= 'F':
		return c - 'A' + 10
	case c >= 'a' && c <= 'f':
		return c - 'a' + 10
	default:
		return 0 // 无效字符返回0
	}
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
