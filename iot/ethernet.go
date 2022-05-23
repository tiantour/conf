package iot

// Ethernet ethernet
type Ethernet struct {
	IP      []string `toml:"ip"`      // 地址
	Rack    int      `toml:"rack"`    // 架子
	Port    int      `toml:"port"`    // 端口
	Slot    int      `toml:"slot"`    // 插槽
	SlaveID byte     `toml:"slaveid"` // 站号
}
