package conf

// Unique unique
type Unique struct {
	Data map[string]unique `toml:"unique"`
}

// NewUnique new unique
func NewUnique() *Unique {
	return &conf.Unique
}

type unique struct {
	Length int    `toml:"length"` // 长度
	Name   string `toml:"name"`   // 名称
}
