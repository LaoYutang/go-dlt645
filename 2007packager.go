package dlt645

import (
	"errors"
	"fmt"
)

type dlt6452007Packager struct{}

func (d *dlt6452007Packager) Encode(pdu IPortocolDataUnit) (adu []byte, err error) {
	return pdu.Value(), nil
}

func (d *dlt6452007Packager) Decode(adu []byte) (pdu IPortocolDataUnit, err error) {
	return NewProtocolDataUnit(adu)
}

type Protocol2007DataUnit struct {
	Front   []byte // 唤醒 在主站发送帧信息之前，先发送1—4个字节0xFE，以唤醒接收方。
	Start   byte   // 起始位 标识一帧信息的开始，其值为68H=01101000B。
	Address []byte // 地址域 地址域由6个字节构成，每字节2位BCD码地址域传输时低字节在前，高字节在后。
	C       byte   // 控制码C
	L       uint8  // 数据域长度L 数据域的字节数。读数据时L≤200，写数据时L≤50，L=0表示无数据域
	Data    []byte // 数据域(未加0x33的数据) 数据域包括数据标识、密码、操作者代码、数据、帧序号等，其结构随控制码的功能而改变。传输时发送方按字节进行加33H处理，接收方按字节进行减33H处理。
	Cs      byte   // 校验码 从第一个帧起始符开始到校验码之前的所有各字节的模256的和，即各字节二进制算术和，不计超过256的溢出值
	End     byte   // 标识一帧信息的结束，其值为16H=00010110B。
}

// 16进制字符串帧
func (a *Protocol2007DataUnit) Value() []byte {
	a.L = uint8(len(a.Data))
	// 计算 CS
	bs := []byte{0x68}              // 起始符
	bs = append(bs, a.Address...)   // 地址域
	bs = append(bs, 0x68, a.C, a.L) // 控制码和数据域长度
	for _, num := range a.Data {
		bs = append(bs, num+0x33) // 数据域
	}
	a.Cs = a.ComputeCs()
	finalBs := []byte{}
	finalBs = append(finalBs, a.Front...)
	finalBs = append(finalBs, bs...)
	finalBs = append(finalBs, a.Cs, a.End)
	return finalBs
}

// 计算校验码
func (p *Protocol2007DataUnit) ComputeCs() byte {
	var sum byte
	sum += p.Start*2 + p.C + p.L
	for _, b := range p.Address {
		sum += b
	}
	for _, b := range p.Data {
		sum += b + 0x33
	}
	return sum
}

// 数据标识
func (p *Protocol2007DataUnit) Identify() string {
	if p.L >= 4 {
		return BcdBytes2String(BytesReverse(p.Data[:4]))
	} else {
		return ""
	}
}

// 根据前一个控制码判断成功失败，返回数据（减去 33H 之后的）
func (p *Protocol2007DataUnit) Result(cmdC byte) (data []byte, err error) {
	errC := cmdC + 0xC0
	if p.C == errC {
		err = fmt.Errorf("response error, code: %#v", p.Data)
	}
	data = p.Data
	return
}

func (p *Protocol2007DataUnit) Verify() bool {
	return p.ComputeCs() == p.Cs
}

// 创建通用协议数据单元 addr-地址(大字序) c-控制码 data-数据(处理好的小字序数据)
func NewCommonProtocolDataUnitByBytes(addr []byte, c byte, data []byte) *Protocol2007DataUnit {
	return &Protocol2007DataUnit{
		Front:   []byte{0xfe, 0xfe, 0xfe, 0xfe},
		Start:   0x68,
		End:     0x16,
		Address: BytesReverse(addr),
		C:       c,
		L:       uint8(len(data)),
		Data:    data,
	}
}

// 解析报文生成协议数据单元
func NewProtocolDataUnit(bs []byte) (*Protocol2007DataUnit, error) {
	pdu := &Protocol2007DataUnit{
		Start: 0x68,
		End:   0x16,
	}
	for _, b := range bs {
		if b == 0xfe {
			pdu.Front = append(pdu.Front, b)
		} else {
			break
		}
	}
	bs = bs[len(pdu.Front):] // 删除前导向
	// 帧起始符 68H
	if bs[0] == 0x68 {
		bs = bs[1:]
		// 地址域
		pdu.Address = BytesReverse(bs[:6])
		bs = bs[6:]
		bs = bs[1:]   // 帧起始符
		pdu.C = bs[0] // 控制码
		pdu.L = bs[1] // 数据域长度
		bs = bs[2:]
		data := bs[:pdu.L] // 数据域
		pdu.Data = BytesSub(data, 0x33)
		bs = bs[pdu.L:]
		pdu.Cs = bs[0]  // 校验码
		pdu.End = bs[1] // 帧结束符

		// 校验失败
		if !pdu.Verify() {
			return nil, errors.New("response invalid checksum")
		}

		return pdu, nil
	} else {
		return nil, errors.New("response invalid frame start")
	}
}
