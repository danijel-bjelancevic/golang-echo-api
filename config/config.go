package config

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)

type Config struct {
	ApiName        string `envconfig:"API_NAME" default:"dolos-risk-service" required:"true"`
	ApiHost        string `envconfig:"API_HOST" default:"localhost" required:"true"`
	ApiPort        int    `envconfig:"API_PORT" default:"8080" required:"true"`
	Debug          bool   `envconfig:"DEBUG" default:"false" required:"true"`
	GrayLogUDPAddr string `envconfig:"GRAYLOG_UDP_ADDR" default:"graylog-udp:12202" required:"false"`
}

//Get will return all configs
func Get() (*Config, error) {
	c := &Config{}

	if err := envconfig.Process("", c); err != nil {
		return nil, errors.Wrap(err, "failed to get configuration from environment")
	}
	return c, nil
}
