package union

// Wechat wechat
type Wechat struct {
	AppID     string `toml:"app_id"`     // 编号
	AppSecret string `toml:"app_secret"` // 密钥
}
