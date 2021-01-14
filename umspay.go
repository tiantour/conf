package conf

// Umspay umspay
type Umspay struct {
	Data map[string]umspay `toml:"umspay"`
}

// NewUmspay new umspay
func NewUmspay() *Umspay {
	return &conf.Umspay
}

type umspay struct {
	AppID  string `toml:"app_id"`  // app_id
	AppKey string `toml:"app_key"` // app_key
	MID    string `toml:"mid"`     // 商户号
	TID    string `toml:"tid"`     // 终端号
}
