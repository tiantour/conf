package conf

// Service service
type Service struct {
	Data map[string]service `toml:"service"`
}

// NewService new service
func NewService() *Service {
	return &conf.Service
}

type service struct {
	Origin string `toml:"origin"` // 域名
	Target string `toml:"target"` // 地址
}
