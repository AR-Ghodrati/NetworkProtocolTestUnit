package main

import (
	"github.com/joho/godotenv"
	"gsm/protocol"
)

func main() {
	godotenv.Load()
	protocol.StartKCP()
}
