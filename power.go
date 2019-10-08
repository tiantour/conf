package conf

// Power power
type Power struct {
	Power []power `toml:"power"`
}

// NewPower new power
func NewPower() *Power {
	return &conf.Power
}

type power struct {
	Name    string  `toml:"name"`    // 备注
	Point   string  `toml:"point"`   // 点位
	Uint    string  `toml:"unit"`    // 单位
	Min     float64 `toml:"min"`     // 最大
	Max     float64 `toml:"max"`     // 最小
	Warning float64 `toml:"warning"` // 报警
	SlaveID byte    `toml:"slaveid"` // 站号
}
