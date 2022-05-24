package iot

// Serial serial
type Serial struct {
	Address  string // 地址
	BaudRate int    // 波特率
	DataBits int    // 数据位
	Parity   string // 校验位
	StopBits int    // 停止位
	SlaveID  []byte // 站号
}
