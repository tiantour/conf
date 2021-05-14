package conf

// Gpio gpio
type Gpio struct {
	Data []*GpioItem `toml:"gpio"`
}

// NewGpio new nats
func NewGpio() *Gpio {
	return &conf.Gpio
}

type GpioItem struct {
	Model  int32  `toml:"model"`  // 型号
	Device string `toml:"device"` // 设备
	Name   string `toml:"name"`   // 名字
	Point  uint8  `toml:"point"`  // 地址
}
