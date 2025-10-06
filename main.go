package main

import (
	"github.com/baturax/Controller/backend"
)

func main() {
	backend.Config()
	backend.HandleAll()
}
