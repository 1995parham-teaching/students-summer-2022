package config

import "github.com/1995parham-teaching/students/internal/db"

type Config struct {
	Debug    bool      `koanf:"debug"`
	Secret   string    `konaf:"secret"`
	Admin    Admin     `koanf:"admin"`
	Database db.Config `koanf:"database"`
}

type Admin struct {
	Username string `koanf:"username"`
	Password string `koanf:"password"`
	Name     string `koanf:"name"`
}
