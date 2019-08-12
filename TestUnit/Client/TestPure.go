package Client

import (
	"gsm/Models"
	"gsm/Utils"
	"log"
	"net"
	"os"
	"time"
)

func RunPureTCP(address string, count uint64, packetStringSize int) {
	log.Println(address)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(1)
	}
	// wait for server to become ready
	time.Sleep(time.Second)
	log.Println("Connecting....")

	defer func() {
		log.Println("Send", count, "Packet DONE!!")
	}()
	var i uint64 = 0
	for i = 0; i < count; i++ {
		data := makeTimestamp()
		random, _ := Utils.GenerateRandomString(packetStringSize)
		_, _ = conn.Write(Utils.Serialize(Models.Message{SequenceNumber: uint64(i), Milis: data, Msg: random, MaxPacketCount: count}))
		time.Sleep(time.Microsecond)
	}
}

func RunPureUDP(address string, count uint64, packetStringSize int) {
	conn, err := net.Dial("udp", address)
	if err != nil {
		log.Println("ERROR", err)
		os.Exit(1)
	}
	// wait for server to become ready
	time.Sleep(time.Second)
	log.Println("Connecting....")

	defer func() {
		log.Println("Send", count, "Packet DONE!!")
	}()
	var i uint64 = 0
	for i = 0; i < count; i++ {
		data := makeTimestamp()
		random, _ := Utils.GenerateRandomString(packetStringSize)
		_, _ = conn.Write(Utils.Serialize(Models.Message{SequenceNumber: uint64(i), Milis: data, Msg: random, MaxPacketCount: count}))
		time.Sleep(time.Microsecond)
	}
}
