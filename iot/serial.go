package iot

// Serial serial
type Serial struct {
	Address  string `mapstructure:"address"`  // 地址
	BaudRate int    `mapstructure:"baudrate"` // 波特率
	DataBits int    `mapstructure:"databits"` // 数据位
	Parity   string `mapstructure:"parity"`   // 校验位
	StopBits int    `mapstructure:"stopbits"` // 停止位
	SlaveID  []byte `mapstructure:"slaveid"`  // 站号
}
