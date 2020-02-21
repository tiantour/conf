package conf

// Gear gear
type Gear struct {
	Data []gear `toml:"gear"`
}

// NewGear new gear
func NewGear() *Gear {
	return &conf.Gear
}

type gear struct {
	Name string `toml:"name"` // 名称
	GPIO int    `toml:"gpio"` // gpio
	Pin  int    `toml:"pin"`  // 针脚
}
