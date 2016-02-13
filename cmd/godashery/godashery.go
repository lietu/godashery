package main

import (
	"log"
	"github.com/lietu/godashery/godashery"
	_ "github.com/lietu/godashery/widgets"
)

func main() {
	log.Printf("Starting up GoDashery...")
	settings := godashery.GetSettings()

	log.Printf("Loading widgets...")

	godashery.LoadWidgets()

	log.Printf("Starting up widgets...")
	go godashery.RunWidgets()

	log.Printf("Starting up server...")
	godashery.RunServer(settings)
	log.Fatalf("Exiting...")
}
