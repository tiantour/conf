package conf

// Token token
type Token struct {
	Secret string `mapstructure:"secret"` // 密钥
	Issuer string `mapstructure:"issuer"` // 发行者
}
