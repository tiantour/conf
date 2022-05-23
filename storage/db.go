package storage

// DB db
type DB struct {
	IP       string `toml:"ip"`       // 地址
	Port     string `toml:"port"`     // 端口
	Database string `toml:"database"` // 数据库
	Uname    string `toml:"uname"`    // 用户名
	Passwd   string `toml:"passwd"`   // 密码
}
