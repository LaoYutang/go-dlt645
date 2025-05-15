package dlt645

import (
	"io"
	"time"

	"github.com/spf13/cast"
)

type dlt645SerialTransporter struct {
	serialPort
}

func (mb *dlt645SerialTransporter) Send(aduRequest []byte) (aduResponse []byte, err error) {
	// 发送自动进行连接
	if err = mb.serialPort.connect(); err != nil {
		return
	}

	// 发送报文
	mb.serialPort.logf("dlt645: sending % x\n", aduRequest)
	if _, err = mb.port.Write(aduRequest); err != nil {
		return
	}
	// 延时在 20ms <= Td <= 500ms
	// 读取报文
	// 1. 先读取 14 个字节 	fefefefe 68 190002031122 68 91 84
	var data [1024]byte
	var n int
	var n1 int

	n, err = ReadAtLeast(mb.port, data[:], 14, 500*time.Millisecond)
	if err != nil {
		return
	}
	// 帧起始符长度
	frontLen := 0
	for _, b := range data {
		if b == 0xfe {
			frontLen++
		} else {
			break
		}
	}
	L := cast.ToInt(data[frontLen+9]) // 数据域长度
	// 总字节数
	bytesToRead := frontLen + 1 + 6 + 1 + 1 + 1 + L + 1 + 1
	// 读取剩余字节
	if n < bytesToRead {
		if bytesToRead > n {
			n1, err = io.ReadFull(mb.port, data[n:bytesToRead])
			n += n1
		}
	}
	aduResponse = data[:n]
	mb.serialPort.logf("dlt645: received % x\n", aduResponse)
	return
}

func (mb *dlt645SerialTransporter) Open() (err error) {
	return mb.Connect()
}

func (mb *dlt645SerialTransporter) Close() (err error) {
	return mb.serialPort.Close()
}
