package main

import (
	"../gsm/Client"
	"../gsm/protocol"
	"fmt"
	"github.com/joho/godotenv"
	"testing"
)

func main() {
	_ = godotenv.Load()
	go protocol.StartKCP()
	Client.Run()

}

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("hello")
	}
}
