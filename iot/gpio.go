package iot

// GPIO gpio
type GPIO struct {
	Model  int32    // 型号
	Device string   // 设备
	Name   []string // 名称
	Point  []uint8  // 地址
	Action uint8    // 地址
}
