package conf

// Serial serial
type Serial struct {
	Data map[string]serial `toml:"serial"`
}

// NewSerial new serial
func NewSerial() *Serial {
	return &conf.Serial
}

type serial struct {
	Address  string `toml:"name"`     // 地址
	BaudRate int    `toml:"baudrate"` // 波特率
	DataBits int    `toml:"databits"` // 数据位
	Parity   string `toml:"parity"`   // 校验位
	StopBits int    `toml:"stopbits"` // 停止位
}
