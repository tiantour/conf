package conf

// Server server
type Server struct {
	Data []server `toml:"server"`
}

// NewServer new server
func NewServer() *Server {
	return &conf.Server
}

type server struct {
	Model   int32   `toml:"model"`   // 型号
	Device  string  `toml:"device"`  // 设备
	Port    string  `toml:"port"`    // 端口
	IP      string  `toml:"ip"`      // 地址
	Uname   string  `toml:"uname"`   // 用户
	Passwd  string  `toml:"passwd"`  // 密码
	Uint    string  `toml:"unit"`    // 单位
	Min     float64 `toml:"min"`     // 最大
	Max     float64 `toml:"max"`     // 最小
	Warning float64 `toml:"warning"` // 报警
}
