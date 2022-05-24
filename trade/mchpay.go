package trade

// Mchpay mchpay
type Mchpay struct {
	ClientCert string `mapstructure:"client_cert"` // 证书
	ClientKey  string `mapstructure:"client_key"`  // 键值
}
