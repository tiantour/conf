package conf

// Template template
type Template struct {
	Data map[string]template `toml:"template"`
}

// NewTemplate new tempalte
func NewTemplate() *Template {
	return &conf.Template
}

type template struct {
	ID   string `toml:"id"`   // 编号
	Name string `toml:"name"` // 名称
}
