package conf

// GPIO gpio
type GPIO struct {
	Data map[string]gpio `toml:"gpio"`
}

// NewGPIO new gpio
func NewGPIO() *GPIO {
	return &conf.GPIO
}

// gpio gpio
type gpio struct {
	Point []uint8 `toml:"point"` // 地址
}
