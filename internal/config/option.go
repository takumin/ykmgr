package config

type Option interface {
	Apply(*Config)
}

type LogLevel string

func (o LogLevel) Apply(c *Config) {
	c.LogLevel = string(o)
}

type ConnectionEndpoint string

func (o ConnectionEndpoint) Apply(c *Config) {
	c.Connection.Endpoint = string(o)
}
