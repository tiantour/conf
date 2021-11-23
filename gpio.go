package conf

// GPIO gpio
type GPIO struct {
	Model  int32    `toml:"model"`  // 型号
	Device string   `toml:"device"` // 设备
	Name   []string `toml:"name"`   // 名称
	Point  []uint8  `toml:"point"`  // 地址
	Action uint8    `toml:"action"` // 地址
}
