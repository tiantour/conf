package conf

// Server server
type Server struct {
	Data server `toml:"server"`
}

// NewServer new server
func NewServer() *Server {
	return &conf.Server
}

type server struct {
	IP   string `toml:"ip"`   // 地址
	Port string `toml:"port"` // 端口
	Host string `toml:"host"` // 域名
	Name string `toml:"name"` // 名称
}
