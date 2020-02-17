package conf

// Fan fan
type Fan struct {
	Data []fan `toml:"fan"`
}

// NewFan new fan
func NewFan() *Fan {
	return &conf.Fan
}

type fan struct {
	Name string `toml:"name"` // 名称
	GPIO int    `toml:"gpio"` // gpio
	Pin  int    `toml:"pin"`  // 针脚
}
