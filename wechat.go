package conf

// Wechat wechat
type Wechat struct {
	Data map[string]wechat `toml:"wechat"`
}

// NewWechat new wechat
func NewWechat() *Wechat {
	return &conf.Wechat
}

type wechat struct {
	AppID     string `toml:"app_id"`     // 编号
	AppSecret string `toml:"app_secret"` // 密钥
}
