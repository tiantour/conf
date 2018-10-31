package conf

// Weibo weibo
type Weibo struct {
	Data weibo `toml:"weibo"`
}

// NewWeibo new weibo
func NewWeibo() *Weibo {
	return &conf.Weibo
}

type weibo struct {
	AppID     string `toml:"app_id"`     // 编号
	AppSecret string `toml:"app_secret"` // 密钥
}
