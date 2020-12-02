package conf

// Gateway gateway
type Gateway struct {
	Data gateway `toml:"gateway"`
}

// NewGateway new gateway
func NewGateway() *Gateway {
	return &conf.Gateway
}

type gateway struct {
	Origin string `toml:"origin"` // 域名
}
