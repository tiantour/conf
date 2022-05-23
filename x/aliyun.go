package x

// Aliyun aliyun
type Aliyun struct {
	AccessKey    string `toml:"access_key_id"`     // 键值
	AccessSecret string `toml:"access_key_secret"` // 密钥
	Sign         string `toml:"sign"`              // 签名
}
