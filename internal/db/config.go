package db

import "time"

type Config struct {
	URL               string        `koanf:"url"`
	Name              string        `koanf:"name"`
	ConnectionTimeout time.Duration `koanf:"connection_timeout"`
}
