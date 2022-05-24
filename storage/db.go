package storage

// DB db
type DB struct {
	Driver string `mapstructure:"dirver"` // dirver
	Source string `mapstructure:"source"` // source
}
