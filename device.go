package conf

// Device device
type Device struct {
	Data []device `toml:"device"`
}

// NewDevice new device
func NewDevice() *Device {
	return &conf.Device
}

type device struct {
	IDC      int32  `toml:"idc"`      // 机房
	Model    int32  `toml:"model"`    // 型号
	Address  string `toml:"address"`  // 地址
	BaudRate int    `toml:"baudrate"` // 波特率
	DataBits int    `toml:"databits"` // 数据位
	Parity   string `toml:"parity"`   // 校验位
	StopBits int    `toml:"stopbits"` // 停止位
	SlaveId  byte   `toml:"slaveid"`  // 站号
}
