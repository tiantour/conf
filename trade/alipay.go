package trade

// Alipay alipay
type Alipay struct {
	AppID      string `toml:"app_id"`      // 编号
	AesKey     string `toml:"aes_key"`     // AES
	PublicKey  string `toml:"public_key"`  // 公钥
	PrivateKey string `toml:"private_key"` // 私钥
}
