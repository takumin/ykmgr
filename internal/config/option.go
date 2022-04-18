package config

type Option interface {
	Apply(*Config)
}

type LogLevel string

func (o LogLevel) Apply(c *Config) {
	c.LogLevel = string(o)
}

type ServerListenURL string

func (o ServerListenURL) Apply(c *Config) {
	c.Server.ListenURL = string(o)
}

type ClientEndpoint string

func (o ClientEndpoint) Apply(c *Config) {
	c.Client.Endpoint = string(o)
}
