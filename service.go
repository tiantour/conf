package conf

// Service service
type Service struct {
	Data service `toml:"service"`
}

// NewService new service
func NewService() *Service {
	return &conf.Service
}

type service struct {
	Sensor string `toml:"sensor"` // 传感器
	Server string `toml:"server"` // 服务器
	Power  string `toml:"power"`  // 动环
	Trade  string `toml:"trade"`  // 交易
}
