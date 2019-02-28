package conf

// CB cb
type CB struct {
	Data cb `toml:"cb"`
}

// NewCB new cb
func NewCB() *CB {
	return &conf.CB
}

type cb struct {
	IP       []string `toml:"ip"`       // 地址
	Port     int32    `toml:"port"`     // 端口
	Keyspace string   `toml:"keyspace"` // 数据库
	Uname    string   `toml:"uname"`    // 用户名
	Passwd   string   `toml:"passwd"`   // 密码
}
