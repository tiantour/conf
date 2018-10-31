package conf

// Dayu dayu
type Dayu struct {
	Data map[string]dayu `toml:"dayu"`
}

// NewDayu new dayu
func NewDayu() *Dayu {
	return &conf.Dayu
}

type dayu struct {
	AppKey    string `toml:"app_key"`    // 键值
	AppSecret string `toml:"app_secret"` // 密钥
	Sign      string `toml:"sign"`       // 签名
}
