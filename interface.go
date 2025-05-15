package dlt645

// 定义日志接口
type ILogger interface {
	Printf(format string, v ...interface{})
}

// 定义数据处理
type IPackager interface {
	Encode(pdu IPortocolDataUnit) (adu []byte, err error)
	Decode(adu []byte) (pdu IPortocolDataUnit, err error)
}

type IPortocolDataUnit interface {
	Value() []byte
	ComputeCs() byte
	Identify() string
	Result(cmdC byte) (data []byte, err error)
}

// 定义传输层
type ITransporter interface {
	Send(aduRequest []byte) (aduResponse []byte, err error)
	Open() (err error)
	Close() (err error)
}

// 整合Packager和Transporter
type IClientHandler interface {
	IPackager
	ITransporter
}
