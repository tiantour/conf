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
	Domain string `toml:"domain"` // 域名
	IP     string `toml:"ip"`     // 地址
	Port   string `toml:"port"`   // 端口
}
