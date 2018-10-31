package conf

// Token token
type Token struct {
	Data token `toml:"token"`
}

// NewToken new token
func NewToken() *Token {
	return &conf.Token
}

type token struct {
	Secret string `toml:"secret"` // 密钥
	Issuer string `toml:"issuer"` // 发行者
}
