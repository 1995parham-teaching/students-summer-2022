package config

import (
	"fmt"
	"log"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/structs"
)

const (
	prefix    = "STUDENTS_"
	delimeter = "."
	seprator  = "__"
)

// STUDENTS_DEBUG -> DEBUG -> debug
// STUDENTS_DATABASE__HOST -> DATABASE__HOST -> database__host -> database.host

func callbackEnv(source string) string {
	base := strings.ToLower(strings.TrimPrefix(source, prefix))

	return strings.ReplaceAll(base, seprator, delimeter)
}

func New() Config {
	k := koanf.New(".")

	// load default configuration from default function
	if err := k.Load(structs.Provider(Default(), "koanf"), nil); err != nil {
		log.Fatalf("error loading default: %s", err)
	}

	// load configuration from file
	if err := k.Load(file.Provider("config.toml"), toml.Parser()); err != nil {
		log.Printf("error loading config.toml: %s", err)
	}

	// load environment variables
	if err := k.Load(env.Provider(prefix, delimeter, callbackEnv), nil); err != nil {
		log.Printf("error loading environment variables: %s", err)
	}

	var instance Config
	if err := k.Unmarshal("", &instance); err != nil {
		log.Fatalf("error unmarshalling config: %s", err)
	}

	fmt.Printf("%+v", instance)

	return instance
}
