package conf

// Qiniu qiniu
type Qiniu struct {
	Data map[string]qiniu `toml:"qiniu"`
}

// NewQiniu new qiniu
func NewQiniu() *Qiniu {
	return &conf.Qiniu
}

type qiniu struct {
	AccessKey string `toml:"access_key"` // 键值
	SecretKey string `toml:"secret_key"` // 密钥
}
