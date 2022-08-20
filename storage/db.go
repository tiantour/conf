package storage

// DB db
type DB struct {
	Driver string `mapstructure:"driver"` // dirver
	Source string `mapstructure:"source"` // source
}
