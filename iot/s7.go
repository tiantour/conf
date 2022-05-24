package iot

type S7 struct {
	Model      int32    `mapstructure:"model"`      // 型号
	Device     string   `mapstructure:"device"`     // 设备
	Port       string   `mapstructure:"port"`       // 端口
	IP         string   `mapstructure:"ip"`         // 地址
	Function   uint16   `mapstructure:"function"`   // 功能
	Location   uint16   `mapstructure:"location"`   // 地址
	Name       []string `mapstructure:"name"`       // 名字
	Address    uint16   `mapstructure:"address"`    // 地址
	Quantity   uint16   `mapstructure:"quantity"`   // 数量
	Amount     uint16   `mapstructure:"amount"`     // 数量
	Coil       uint16   `mapstructure:"coil"`       // 线圈
	Inverse    uint16   `mapstructure:"inverse"`    // 反码
	Complement uint16   `mapstructure:"complement"` // 补码
	Integer    uint16   `mapstructure:"integer"`    // 整数
	Decimal    uint16   `mapstructure:"decimal"`    // 小数
	Plus       float64  `mapstructure:"plus"`       // 加法
	Times      float64  `mapstructure:"times"`      // 乘法
	Uint       string   `mapstructure:"unit"`       // 单位
}
