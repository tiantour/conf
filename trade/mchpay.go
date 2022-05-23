package trade

// Mchpay mchpay
type Mchpay struct {
	ClientCert string `toml:"client_cert"` // 证书
	ClientKey  string `toml:"client_key"`  // 键值
}
