package iot

// IPMI ipmi
type IPMI struct {
	Model  int32  `mapstructure:"model"`  // 型号
	Device string `mapstructure:"device"` // 设备
	Port   string `mapstructure:"port"`   // 端口
	IP     string `mapstructure:"ip"`     // 地址
	Uname  string `mapstructure:"uname"`  // 用户
	Passwd string `mapstructure:"passwd"` // 密码
	Uint   string `mapstructure:"unit"`   // 单位
}
