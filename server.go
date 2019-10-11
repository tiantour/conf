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
	IDC       int32   `toml:"idc"`       // 机房
	Model     int32   `toml:"model"`     // 型号
	Name      string  `toml:"name"`      // 备注
	Address   string  `toml:"address"`   // 链接
	Uname     string  `toml:"uname"`     // 用户
	Passwd    string  `toml:"passwd"`    // 密码
	Community string  `toml:"community"` // 社区
	Point     string  `toml:"point"`     // 位置
	Uint      string  `toml:"unit"`      // 单位
	Min       float64 `toml:"min"`       // 最大
	Max       float64 `toml:"max"`       // 最小
	Warning   float64 `toml:"warning"`   // 报警
}
