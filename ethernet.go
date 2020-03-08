package conf

// Ethernet ethernet
type Ethernet struct {
	Data map[string]ethernet `toml:"ethernet"`
}

// NewEthernet new ethernet
func NewEthernet() *Ethernet {
	return &conf.Ethernet
}

type ethernet struct {
	IP string `toml:"ip"` // 地址
}
