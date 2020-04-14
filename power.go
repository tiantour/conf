package conf

// Power power
type Power struct {
	Data []power `toml:"power"`
}

// NewPower new power
func NewPower() *Power {
	return &conf.Power
}

type power struct {
	Model    int32    `toml:"model"`    // 型号
	Device   string   `toml:"device"`   // 设备
	Port     string   `toml:"port"`     // 端口
	SlaveID  byte     `toml:"slaveid"`  // 站号
	Function uint16   `toml:"function"` // 功能
	Name     []string `toml:"name"`     // 名字
	Address  uint16   `toml:"address"`  // 地址
	Quantity uint16   `toml:"quantity"` // 数量
	Amount   uint16   `toml:"amount"`   // 数量
	Reverse  uint16   `toml:"reverse"`  // 反转
	Signed   uint16   `toml:"signed"`   // 符号
	Integer  uint16   `toml:"integer"`  // 整数
	Decimal  uint16   `toml:"decimal"`  // 小数
	Switch   uint16   `toml:"switch"`   // 开关
	Uint     string   `toml:"unit"`     // 单位
	Min      float64  `toml:"min"`      // 最大
	Max      float64  `toml:"max"`      // 最小
	Warning  float64  `toml:"warning"`  // 报警
	Compute  string   `toml:"compute"`  // 计算
}
