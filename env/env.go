package env

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v2"
)

func NewAppEnv() Env {
	return NewEnv(false)
}

func NewTestEnv() Env {
	return NewEnv(true)
}

func NewEnv(t bool) Env {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}
	stage := os.Getenv("STAGE")

	buffer, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		log.Fatalln(err)
	}

	s := &settings{}

	err = yaml.Unmarshal(buffer, &s)
	if err != nil {
		log.Fatalln(err)
	}

	if t {
		return s.Test
	}

	var e Env

	switch stage {
	case "development":
		e = s.Development
	case "production":
		e = s.Production
	case "test":
		e = s.Test
	default:
		e = s.Development
	}

	e.Stage = stage
	return e
}

type Env struct {
	Stage string
	Dsn   string `yaml:"datasource"`
}

type settings struct {
	Test        Env
	Development Env
	Production  Env
}
