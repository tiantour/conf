package conf

// Mchpay mchpay
type Mchpay struct {
	Data mchpay `toml:"mchpay"`
}

// NewMchpay new mchpay
func NewMchpay() *Mchpay {
	return &conf.Mchpay
}

type mchpay struct {
	ClientCert string `toml:"client_cert"` // 证书
	ClientKey  string `toml:"client_key"`  // 键值
}
