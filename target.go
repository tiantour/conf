package conf

// Target target
type Target struct {
	Data []target `toml:"target"`
}

// NewTarget new target
func NewTarget() *Target {
	return &conf.Target
}

type target struct {
	Model     int32   `toml:"model"`     // 型号
	Device    string  `toml:"device"`    // 设备
	Port      string  `toml:"port"`      // 端口
	IP        string  `toml:"ip"`        // 地址
	Community string  `toml:"community"` // 社区
	Min       float64 `toml:"min"`       // 最大
	Max       float64 `toml:"max"`       // 最小
	Uint      string  `toml:"unit"`      // 单位
	Notice    string  `toml:"notice"`    // 通知
	Compute   string  `toml:"compute"`   // 计算
}
