package conf

// Aliyun aliyun
type Aliyun struct {
	Data map[string]aliyun `toml:"aliyun"`
}

// NewAliyun new aliyun
func NewAliyun() *Aliyun {
	return &conf.Aliyun
}

type aliyun struct {
	AccessKey    string `toml:"access_key_id"`     // 键值
	AccessSecret string `toml:"access_key_secret"` // 密钥
	Sign         string `toml:"sign"`              // 签名
}
