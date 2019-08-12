package main

import (
	"bufio"
	"fmt"
	"gsm/TestUnit"
	"log"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 4096)

	fmt.Print("Run As 1-Client Or 2-Server : ")
	for {
		char, _, _ := reader.ReadRune()
		switch char {
		case '1':
			for {
				reader := bufio.NewReaderSize(os.Stdin, 4096)

				fmt.Println("Select Client Type : ")
				fmt.Println("1-QUIC")
				fmt.Println("2-KCP")
				fmt.Println("3-Pure TCP")
				fmt.Println("4-Pure UDP")

				var _type string
				_, _ = fmt.Scanf("%s", &_type)

				fmt.Print("Enter address (IP:PORT):")
				var address string
				_, _ = fmt.Scanf("%s", &address)

				fmt.Print("Enter Packet Count :")
				line, _, _ := reader.ReadLine()
				packetCount, _ := strconv.ParseInt(string(line), 10, 64)

				fmt.Print("Enter Packet Size (Bytes) :")
				_line, _, _ := reader.ReadLine()
				PS, _ := strconv.ParseInt(string(_line), 10, 16)

				switch _type {
				case "1":
					TestUnit.RunClientQUIC(address, uint64(packetCount), int(PS))
					break
				case "2":
					TestUnit.RunClientKCP(address, uint64(packetCount), int(PS))
					break
				case "3":
					TestUnit.RunClientTCP(address, uint64(packetCount), int(PS))
					break
				case "4":
					TestUnit.RunClientUDP(address, uint64(packetCount), int(PS))
					break
				default:
					log.Println("Select Valid Option")
				}
			}
		case '2':
			var PORT string
			fmt.Print("Select Server PORT : ")
			_, _ = fmt.Scanf("%s", &PORT)

			fmt.Println("Select Server Type : ")
			fmt.Println("1-QUIC")
			fmt.Println("2-KCP")
			fmt.Println("3-Pure TCP")
			fmt.Println("4-Pure UDP")

			for {
				var _type string
				_, _ = fmt.Scanf("%s", &_type)

				switch _type {
				case "1":
					TestUnit.RunServerQUIC(PORT)
					break
				case "2":
					TestUnit.RunServerKCP(PORT)
					break
				case "3":
					TestUnit.RunServerPureTCP(PORT)
					break
				case "4":
					TestUnit.RunServerPureUDP(PORT)
					break
				default:
					log.Println("Select Valid Option")
				}
			}
		}
	}
}
