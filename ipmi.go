package conf

// Ipmi ipmi
type Ipmi struct {
	Data []*IpmiItem `toml:"ipmi"`
}

// NewIpmi new ipmi
func NewIpmi() *Ipmi {
	return &conf.Ipmi
}

type IpmiItem struct {
	Model  int32  `toml:"model"`  // 型号
	Device string `toml:"device"` // 设备
	Port   string `toml:"port"`   // 端口
	IP     string `toml:"ip"`     // 地址
	Uname  string `toml:"uname"`  // 用户
	Passwd string `toml:"passwd"` // 密码
	Uint   string `toml:"unit"`   // 单位
}
