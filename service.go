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
	Liquid string `toml:"liquid"` // 液冷 10
	Power  string `toml:"power"`  // 动环 20
	Sensor string `toml:"sensor"` // 传感器 30
	Server string `toml:"server"` // 服务器 40
	Trade  string `toml:"trade"`  // 交易
}
