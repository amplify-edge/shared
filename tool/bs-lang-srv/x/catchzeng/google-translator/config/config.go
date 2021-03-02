package config

import (
	"fmt"

	"github.com/timest/env"
)

//Env Env
var Env *envs

type envs struct {
	GoogleURL string `env:"GOOGLE_URL" default:"https://translate.google.cn"`
	ChunkSize int    `env:"CHUNK_SIZE" default:"5000"`
}

func init() {
	Env = new(envs)
	env.IgnorePrefix()
	err := env.Fill(Env)
	fmt.Println(Env)
	if err != nil {
		panic(err)
	}
}
