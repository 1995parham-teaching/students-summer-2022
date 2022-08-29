package config

import "github.com/1995parham-teaching/students/internal/db"

type Config struct {
	Debug    bool      `koanf:"debug"`
	Database db.Config `koanf:"database"`
}
