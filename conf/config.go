package conf

type Config struct {
	Enabled      bool
	DatabasePath string
	Port         string
}

func NewConfig() *Config {
	return &Config{
		Enabled:      true,
		DatabasePath: "root:123456@tcp(127.0.0.1:3306)/tiki_sync",
		Port:         "8000",
	}
}
