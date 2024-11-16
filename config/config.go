package config

type Config struct {
	Width  int
	Height int
}

func New() *Config {
	deff := &Config{
		Width:  100,
		Height: 50,
	}
	return deff
}
