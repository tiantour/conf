package conf

// IDC IDC
type IDC struct {
	Data idc `toml:"idc"`
}

// NewIDC new idc
func NewIDC() *IDC {
	return &conf.IDC
}

// idc idc
type idc struct {
	Number int32  `toml:"number"` // 编号
	Name   string `toml:"name"`   // 名称
}
