package port

// GPIO gpio
type GPIO struct {
	Model  int32    `mapstructure:"model"`  // 型号
	Device string   `mapstructure:"device"` // 设备
	Name   []string `mapstructure:"name"`   // 名称
	Point  []uint8  `mapstructure:"point"`  // 地址
	Action uint8    `mapstructure:"action"` // 地址
}
