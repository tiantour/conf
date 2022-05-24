package iot

// IPMI ipmi
type IPMI struct {
	Model  int32  // 型号
	Device string // 设备
	Port   string // 端口
	IP     string // 地址
	Uname  string // 用户
	Passwd string // 密码
	Uint   string // 单位
}
