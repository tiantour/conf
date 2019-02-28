package conf

// Mafengwo mafengwo
type Mafengwo struct {
	Data map[string]mafengwo `toml:"mafengwo"`
}

// NewMafengwo new mafengwo
func NewMafengwo() *Mafengwo {
	return &conf.Mafengwo
}

type mafengwo struct {
	ClientID     int32  `toml:"client_id"`     // 客户端编号
	PartnerID    int32  `toml:"partner_id"`    // 合作方编号
	ClientSecret string `toml:"client_secret"` // 客户端密钥
	AseKey       string `toml:"ase_key"`       // 加密密钥
}
