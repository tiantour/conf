package trade

// Wxpay wxpay
type Wxpay struct {
	MchID      string `toml:"mch_id"`      // 编号
	Key        string `toml:"key"`         // API密钥
	Key3       string `toml:"key3"`        // API密钥v3
	ClientCert string `toml:"client_cert"` // 证书
	ClientKey  string `toml:"client_key"`  // 私钥
	SerialNo   string `toml:"serial_no"`   // 证书编号
}
