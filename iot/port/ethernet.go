package port

// Ethernet ethernet
type Ethernet struct {
	IP      []string `mapstructure:"ip"`      // 地址
	Rack    int      `mapstructure:"rack"`    // 架子
	Port    int      `mapstructure:"port"`    // 端口
	Slot    int      `mapstructure:"slot"`    // 插槽
	SlaveID byte     `mapstructure:"slaveid"` // 站号
}
