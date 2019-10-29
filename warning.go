package conf

// Warning warning
type Warning struct {
	Data []warning `toml:"warning"`
}

// NewWarning new warning
func NewWarning() *Warning {
	return &conf.Warning
}

type warning struct {
	SlaveID  byte    `toml:"slaveid"`  // 站号
	Function uint16  `toml:"function"` // 功能
	IP       string  `toml:"ip"`       // 地址
	Location uint16  `toml:"location"` // 地址
	Name     string  `toml:"name"`     // 名字
	Address  uint16  `toml:"address"`  // 地址
	Quantity uint16  `toml:"quantity"` // 数量
	Reverse  uint16  `toml:"reverse"`  // 反转
	Decimal  uint16  `toml:"decimal"`  // 小数
	Uint     string  `toml:"unit"`     // 单位
	Min      float64 `toml:"min"`      // 最大
	Max      float64 `toml:"max"`      // 最小
	Warning  float64 `toml:"warning"`  // 报警
}
