package config

type Config struct {
	LogLevel string

	Connection struct {
		Endpoint string
	}
}

func NewConfig(opts ...Option) *Config {
	c := &Config{}
	for _, o := range opts {
		o.Apply(c)
	}
	return c
}
