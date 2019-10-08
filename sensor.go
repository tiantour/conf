package conf

// Sensor sensor
type Sensor struct {
	Data []sensor `toml:"sensor"`
}

// NewSensor new sensor
func NewSensor() *Sensor {
	return &conf.Sensor
}

// sensor sensor
type sensor struct {
	Name    string  `toml:"name"`    // 名称
	Point   string  `toml:"point"`   // 地址
	Uint    string  `toml:"unit"`    // 单位
	Min     float64 `toml:"min"`     // 最大
	Max     float64 `toml:"max"`     // 最小
	Warning float64 `toml:"warning"` // 报警
	SlaveID byte    `toml:"slaveid"` // 站号
}
