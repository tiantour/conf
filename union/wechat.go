package union

// Wechat wechat
type Wechat struct {
	AppID     string `mapstructure:"app_id"`     // 编号
	AppSecret string `mapstructure:"app_secret"` // 密钥
}
