package conf

// Nats nats
type Nats struct {
	Data nats `toml:"nats"`
}

// NewNats new cache
func NewNats() *Nats {
	return &conf.Nats
}

type nats struct {
	IP   string `toml:"ip"`   // 地址
	Port string `toml:"port"` // 端口
}
