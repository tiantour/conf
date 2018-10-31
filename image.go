package conf

// Image image
type Image struct {
	Data map[string]image `toml:"image"`
}

// NewImage new image
func NewImage() *Image {
	return &conf.Image
}

type image struct {
	Bucket string `toml:"bucket"` // 空间
	Host   string `toml:"host"`   // 域名
	Uname  string `toml:"uname"`  // 用户名
	Passwd string `toml:"passwd"` // 密码
}
