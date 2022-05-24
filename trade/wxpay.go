package trade

// Wxpay wxpay
type Wxpay struct {
	MchID      string `mapstructure:"mch_id"`      // 编号
	Key        string `mapstructure:"key"`         // API密钥
	Key3       string `mapstructure:"key3"`        // API密钥v3
	ClientCert string `mapstructure:"client_cert"` // 证书
	ClientKey  string `mapstructure:"client_key"`  // 私钥
	SerialNo   string `mapstructure:"serial_no"`   // 证书编号
}
