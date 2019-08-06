package TestUnit

import (
	"../TestUnit/Client"
	"github.com/joho/godotenv"
)

func RunClient() {
	_ = godotenv.Load()
	Client.Run(Count)
}
