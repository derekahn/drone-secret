package main

import (
	"context"
	"log"

	"github.com/derekahn/drone-secret/plugin"
	"github.com/kelseyhightower/envconfig"
)

func main() {
	log.Println("🤫")

	var args plugin.Args
	if err := envconfig.Process("", &args); err != nil {
		log.Fatalln(err)
	}

	if err := plugin.Exec(context.Background(), args); err != nil {
		log.Fatalln(err)
	}
	log.Println("🚀")
}
