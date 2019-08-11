package config

type DatabaseConfig struct {
	Type         string `default:"sqlite3"`
	File         string `default:"/tmp/huigo.db"`
	MaxIdleConns int    `default:"5"`
	MaxOpenConns int    `default:"5"`
}

type Config struct {
	DatabaseConfig
	IsDevelopment bool `default:"True"`
}