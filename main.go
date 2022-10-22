package main

import (
	"log"

	"github.com/pocketbase/pocketbase"

	_ "pocketbase-leave-management/migrations"
)

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
