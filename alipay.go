package conf

// Alipay alipay
type Alipay struct {
	Data alipay `toml:"alipay"`
}

// NewAlipay new alipay
func NewAlipay() *Alipay {
	return &conf.Alipay
}

type alipay struct {
	AppID      string `toml:"app_id"`      // 编号
	PublicKey  string `toml:"public_key"`  // 公钥
	PrivateKey string `toml:"private_key"` // 私钥
}
