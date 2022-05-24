package iot

// Ethernet ethernet
type Ethernet struct {
	IP      []string // 地址
	Rack    int      // 架子
	Port    int      // 端口
	Slot    int      // 插槽
	SlaveID byte     // 站号
}
