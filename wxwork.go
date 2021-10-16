package conf

// Wxwork wxwork
type Wxwork struct {
	Data map[string]wxwork `toml:"wxwork"`
}

// NewWxwork new wxwork
func NewWxwork() *Wxwork {
	return &conf.Wxwork
}

type wxwork struct {
	CorpID     string `toml:"corp_id"`     // 编号
	CorpSecret string `toml:"corp_secret"` // 密钥
}
