package conf

// Light light
type Light struct {
	Data []light `toml:"light"`
}

// NewLight new fan
func NewLight() *Light {
	return &conf.Light
}

type light struct {
	Name string `toml:"name"` // 名称
	GPIO int    `toml:"gpio"` // gpio
	Pin  int    `toml:"pin"`  // 针脚
}
