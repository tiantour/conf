package x

// Aliyun aliyun
type Aliyun struct {
	AccessKey    string `mapstructure:"access_key_id"`     // 键值
	AccessSecret string `mapstructure:"access_key_secret"` // 密钥
	Sign         string `mapstructure:"sign"`              // 签名
}
