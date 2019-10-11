package conf

// Point point
type Point struct {
	Data []point `toml:"point"`
}

// NewPoint new point
func NewPoint() *Point {
	return &conf.Point
}

type point struct {
	Name    string  `toml:"name"`    // 备注
	Address string  `toml:"address"` // 地址
	Uint    string  `toml:"unit"`    // 单位
	Min     float64 `toml:"min"`     // 最大
	Max     float64 `toml:"max"`     // 最小
	Warning float64 `toml:"warning"` // 报警
}
