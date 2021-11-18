package conf

// DO do
type DO struct {
	Model  int32  `toml:"model"`  // 型号
	Device string `toml:"device"` // 设备
	Point  uint8  `toml:"point"`  // 地址
	Action uint8  `toml:"action"` // 地址
}
