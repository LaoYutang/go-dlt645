package dlt645

import "time"

type serial2007Handler struct {
	dlt6452007Packager
	dlt645SerialTransporter
}

// 创建一个串口通讯下的2007协议处理器
func NewSerial2007Handler(address string) *serial2007Handler {
	handler := &serial2007Handler{}
	handler.Name = address
	handler.ReadTimeout = time.Millisecond * 500
	return handler
}
