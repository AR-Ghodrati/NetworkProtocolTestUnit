package main

import (
	"bufio"
	"fmt"
	"gsm/TestUnit"
	"log"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select Server PORT : ")
	line, _, _ := reader.ReadLine()
	PORT := string(line)

	fmt.Println("Select Server Type : ")
	fmt.Println("1-QUIC")
	fmt.Println("2-KCP")

	for {
		char, _, _ := reader.ReadRune()
		switch char {
		case '1':
			TestUnit.RunServerQUIC(PORT)
			break
		case '2':
			TestUnit.RunServerKCP(PORT)
			break
		default:
			log.Println("Select Valid Option")
		}
	}

}
