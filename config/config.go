package config

// Config defines the structure of the application's configuration settings.
type Config struct {
	TGBotHost   string `toml:"tg_bot_host"`
	StoragePath string `toml:"storage_path"`
	BatchSize   int    `toml:"batch_size"`
	DatabaseURL string `toml:"database_url"`
	TGToken     string `toml:"telegram_token"`
}

// NewConfig returns a default configuration with pre-defined values.
func NewConfig() *Config {
	return &Config{
		TGBotHost: "api.telegram.org",
	}
}
