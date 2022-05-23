package conf

// Token token
type Token struct {
	Secret string `toml:"secret"` // 密钥
	Issuer string `toml:"issuer"` // 发行者
}
