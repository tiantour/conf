package conf

// GPIO gpio
type GPIO struct {
	Data map[string]gpio `toml:"gpio"`
}

// NewGPIO new gpio
func NewGPIO() *GPIO {
	return &conf.GPIO
}

type gpio struct {
	Name string `toml:"name"` // 名称
	GPIO int    `toml:"gpio"` // gpio
	Pin  int    `toml:"pin"`  // 针脚
}
