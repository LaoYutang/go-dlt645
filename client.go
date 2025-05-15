package dlt645

import "errors"

type Dlt645Client struct {
	packager    IPackager
	transporter ITransporter
}

// NewClient 创建一个新的客户端
func NewClient(handler IClientHandler) *Dlt645Client {
	return &Dlt645Client{packager: handler, transporter: handler}
}

// 读取数据 addr-地址 id-数据标识符
func (c *Dlt645Client) ReadData(addr string, id string) (results []byte, err error) {
	pdu := NewCommonProtocolDataUnitByBytes(
		String2BcdBytes(addr), 0x11,
		BytesReverse(String2BcdBytes(id)),
	)
	res, err := c.send(pdu)
	if err != nil {
		return
	}
	results, err = res.Result(pdu.C)
	if err != nil {
		return
	}
	return
}

// 设置参数 addr-地址 id-数据标识符 data-数据(未加33H)
func (c *Dlt645Client) SetParam(addr string, id string, data []byte) (results []byte, err error) {
	pdu := NewCommonProtocolDataUnitByBytes(
		String2BcdBytes(addr), 0x14,
		append(BytesReverse(String2BcdBytes(id)), data...),
	)
	res, err := c.send(pdu)
	if err != nil {
		return
	}
	results, err = res.Result(pdu.C)
	if err != nil {
		return
	}
	return
}

// 发送自定义pdu
func (c *Dlt645Client) Send(pduSource any) (results []byte, err error) {
	pdu, ok := pduSource.(*Protocol2007DataUnit)
	if !ok {
		return nil, errors.New("invalid Protocol2007DataUnit")
	}

	res, err := c.send(pdu)
	if err != nil {
		return
	}
	results, err = res.Result(pdu.C)
	if err != nil {
		return
	}
	return
}

func (c *Dlt645Client) Open() (err error) {
	err = c.transporter.Open()
	return
}

func (c *Dlt645Client) Close() (err error) {
	err = c.transporter.Close()
	return
}

func (c *Dlt645Client) send(request *Protocol2007DataUnit) (response IPortocolDataUnit, err error) {
	aduRequest, err := c.packager.Encode(request)
	if err != nil {
		return
	}
	aduResponse, err := c.transporter.Send(aduRequest)
	if err != nil {
		return
	}
	pdu, err := c.packager.Decode(aduResponse)
	if err != nil {
		return
	}
	return pdu, nil
}
