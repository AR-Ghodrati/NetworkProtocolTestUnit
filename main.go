package main

import (
	"../gsm/protocol"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	protocol.StartQUIC()
}
