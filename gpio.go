package conf

// GPIO gpio
type GPIO struct {
	Model  int32  `toml:"model"`  // 型号
	Device string `toml:"device"` // 设备
	Name   string `toml:"name"`   // 名字
	Point  uint8  `toml:"point"`  // 地址
}
