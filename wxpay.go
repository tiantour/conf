package conf

// Wxpay wxpay
type Wxpay struct {
	Data map[string]wxpay `toml:"wxpay"`
}

// NewWxpay new wxpay
func NewWxpay() *Wxpay {
	return &conf.Wxpay
}

type wxpay struct {
	MchID string `toml:"mch_id"` // 编号
	Key   string `toml:"key"`    // 键值
}
