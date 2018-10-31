package conf

// QQ qq
type QQ struct {
	Data qq `toml:"qq"`
}

// NewQQ new qq
func NewQQ() *QQ {
	return &conf.QQ
}

type qq struct {
	AppID  string `toml:"app_id"`  // 编号
	AppKey string `toml:"app_key"` // 键值
}
