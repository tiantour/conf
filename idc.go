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
	Interval int32    `toml:"interval"` // 间隔
	Number   int32    `toml:"number"`   // 编号
	Name     string   `toml:"name"`     // 名称
	Template string   `toml:"template"` // 模板
	Phone    []string `toml:"phone"`    // 电话
}
