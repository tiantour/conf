package trade

// Alipay alipay
type Alipay struct {
	AppID      string `mapstructure:"app_id"`      // 编号
	AesKey     string `mapstructure:"aes_key"`     // AES
	PublicKey  string `mapstructure:"public_key"`  // 公钥
	PrivateKey string `mapstructure:"private_key"` // 私钥
}
