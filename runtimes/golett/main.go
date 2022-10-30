package main

import (
	"log"

	"github.com/spf13/pflag"

	"github.com/airconduct/kruntime/runtimes/golett/app"
)

func main() {
	pflag.CommandLine = pflag.NewFlagSet("golett", pflag.ExitOnError)
	if err := app.New().Execute(); err != nil {
		log.Fatal(err)
	}
}
