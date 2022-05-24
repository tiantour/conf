package x

// Mafengwo mafengwo
type Mafengwo struct {
	ClientID     int32  `mapstructure:"client_id"`     // 客户端编号
	PartnerID    int32  `mapstructure:"partner_id"`    // 合作方编号
	ClientSecret string `mapstructure:"client_secret"` // 客户端密钥
	AseKey       string `mapstructure:"ase_key"`       // 加密密钥
}
