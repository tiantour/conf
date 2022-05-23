package trade

// Umspay umspay
type Umspay struct {
	AppID  string `toml:"app_id"`  // appid
	AppKey string `toml:"app_key"` // appkey
	MID    string `toml:"mid"`     // 商户号
	TID    string `toml:"tid"`     // 终端号
	Key    string `toml:"key"`     // 密钥
}
