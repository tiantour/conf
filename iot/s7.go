package iot

type S7 struct {
	Model      int32    // 型号
	Device     string   // 设备
	Port       string   // 端口
	IP         string   // 地址
	Function   uint16   // 功能
	Location   uint16   // 地址
	Name       []string // 名字
	Address    uint16   // 地址
	Quantity   uint16   // 数量
	Amount     uint16   // 数量
	Coil       uint16   // 线圈
	Inverse    uint16   // 反码
	Complement uint16   // 补码
	Integer    uint16   // 整数
	Decimal    uint16   // 小数
	Plus       float64  // 加法
	Times      float64  // 乘法
	Uint       string   // 单位
}
