# GO-DLT645
DLT645客户端(主站)的golang实现

## 当前实现
- [x] DLT645-2007协议
- [ ] DLT645-1997协议
- [x] 串口通讯

## 快速开始
```go
// 创建串口2007客户端
handler := dlt645.NewSerial2007Handler("COM13")
handler.Baud = 2400
handler.StopBits = 1
handler.Parity = serial.ParityEven
handler.Size = 8
handler.ReadTimeout = 1 * time.Second
handler.Logger = log.New(os.Stdout, "rs485: ", log.LstdFlags)
client := dlt645.NewClient(handler)
// 读取数据
r, err := client.ReadData("000000000010", "02020100")
result := dlt645.BcdBytes2Float64(r[4:], 3, false)
// 设置数据
r, err := client.SetParam(
  "000000000010", "04000104",
  []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, // 密码、用户代码、数据需要自行反序
)
// 自定义发送
pdu := dlt645.NewCommonProtocolDataUnitByBytes(
  []byte{0x99, 0x99, 0x99, 0x99, 0x99, 0x99}, 0x13,
  []byte{},
)
r, err := client.Send(pdu)
```

## License
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

