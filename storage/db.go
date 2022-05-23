package storage

// DB db
type DB struct {
	Driver string `toml:"dirver"` // dirver
	Source string `toml:"source"` // source
}
