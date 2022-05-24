package trade

// Umspay umspay
type Umspay struct {
	AppID  string `mapstructure:"app_id"`  // appid
	AppKey string `mapstructure:"app_key"` // appkey
	MID    string `mapstructure:"mid"`     // 商户号
	TID    string `mapstructure:"tid"`     // 终端号
	Key    string `mapstructure:"key"`     // 密钥
}
