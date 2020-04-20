package conf

// Next next
type Next struct {
	Data []next `toml:"next"`
}

// NewNext new next
func NewNext() *Next {
	return &conf.Next
}

// next next
type next struct {
	Model      int32    `toml:"model"`      // 型号
	Device     string   `toml:"device"`     // 设备
	Port       string   `toml:"port"`       // 端口
	IP         string   `toml:"ip"`         // 地址
	Function   uint16   `toml:"function"`   // 功能
	Location   uint16   `toml:"location"`   // 地址
	Name       []string `toml:"name"`       // 名字
	Address    uint16   `toml:"address"`    // 地址
	Quantity   uint16   `toml:"quantity"`   // 数量
	Amount     uint16   `toml:"amount"`     // 数量
	Coil       uint16   `toml:"coil"`       // 线圈
	Inverse    uint16   `toml:"inverse"`    // 反码
	Complement uint16   `toml:"complement"` // 补码
	Integer    uint16   `toml:"integer"`    // 整数
	Decimal    uint16   `toml:"decimal"`    // 小数
	Plus       float64  `toml:"plus"`       // 加法
	Times      float64  `toml:"times"`      // 乘法
	Min        float64  `toml:"min"`        // 最大
	Max        float64  `toml:"max"`        // 最小
	Uint       string   `toml:"unit"`       // 单位
	Notice     string   `toml:"notice"`     // 通知
	Compute    string   `toml:"compute"`    // 计算
}
