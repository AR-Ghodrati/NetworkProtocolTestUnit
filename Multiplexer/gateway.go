package Multiplexer

import (
	"context"
	"fmt"
	"github.com/lucas-clemente/quic-go"
	"github.com/xtaci/kcp-go"
	"gsm/Utils"
	"log"
	"net"
	"os"
	"sort"
	"time"
)

const defaultBufferSize = uint(16384)

var NewLine = fmt.Sprintf("\r\n")

func MultiplexPureTCP(conn net.Conn) {
	t1 := makeTimestamp()
	_ = os.Mkdir("Logs", os.ModePerm)
	f, err := os.OpenFile("Logs/"+conn.LocalAddr().String()+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("File Not Opened!!")
	}

	buf := make([]byte, defaultBufferSize)
	var totalP uint64 = 0
	var totalSize uint64 = 0
	var PSize uint64 = 0

	//map [time Diff] = Count
	var DiffMap = make(map[int64]uint64)
	defer func() {
		t2 := makeTimestamp()
		f.WriteString(NewLine)

		// Sort Result
		keys := make([]int64, 0, len(DiffMap))
		for k := range DiffMap {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, k := range keys {
			txt := fmt.Sprintf("With %d Milliseconds Different -> Count is: %d \r\n", k, DiffMap[k])
			f.WriteString(txt)
		}
		log.Println("DONE")
		f.WriteString(NewLine)
		f.WriteString(NewLine)
		f.WriteString("Total Packet Receive : " + fmt.Sprintf("%d", totalP))
		f.WriteString(NewLine)
		f.WriteString("Packet Size : " + fmt.Sprintf("%d Byte", PSize))
		f.WriteString(NewLine)
		f.WriteString("Total Data Byte Receive: " + fmt.Sprintf("%d Bytes ", totalSize))
		f.WriteString(NewLine)
		f.WriteString("Protocol : Pure TCP")
		f.WriteString(NewLine)
		f.WriteString("Stop Receiving Time : " + time.Now().String())
		f.WriteString(NewLine)
		f.WriteString("Time Escaped To Receiving : " + fmt.Sprintf("%d", t2-t1) + " Milliseconds")
		f.WriteString(NewLine)
		f.WriteString("==============================")
		f.Close()
	}()

	f.WriteString(NewLine)
	f.WriteString(NewLine)
	f.WriteString(NewLine)

	f.WriteString("==============================")
	f.WriteString(NewLine)
	f.WriteString("Receiving Data From " + conn.LocalAddr().String() + NewLine)
	f.WriteString("Start Receiving Time : " + time.Now().String())
	f.WriteString(NewLine)

	for {
		lengthR, err := conn.Read(buf)
		msg := Utils.Deserialize(buf[:lengthR])
		i := makeTimestamp()

		totalSize += uint64(lengthR)

		if totalP == 0 {
			totalP = msg.MaxPacketCount
		}
		if PSize == 0 {
			PSize = uint64(lengthR)
		}

		DiffMap[i-msg.Milis] = DiffMap[i-msg.Milis] + 1

		if err != nil {
			log.Fatal(err)
		}
		if msg.SequenceNumber == msg.MaxPacketCount-1 {
			break
		}
	}
}

func MultiplexPureUDP(conn net.Conn) {
	t1 := makeTimestamp()
	_ = os.Mkdir("Logs", os.ModePerm)
	f, err := os.OpenFile("Logs/"+conn.LocalAddr().String()+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("File Not Opened!!")
	}

	buf := make([]byte, defaultBufferSize)
	var totalP uint64 = 0
	var totalSize uint64 = 0
	var PSize uint64 = 0

	//map [time Diff] = Count
	var DiffMap = make(map[int64]uint64)
	defer func() {
		t2 := makeTimestamp()
		f.WriteString(NewLine)

		// Sort Result
		keys := make([]int64, 0, len(DiffMap))
		for k := range DiffMap {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, k := range keys {
			txt := fmt.Sprintf("With %d Milliseconds Different -> Count is: %d \r\n", k, DiffMap[k])
			f.WriteString(txt)
		}
		log.Println("DONE")
		f.WriteString(NewLine)
		f.WriteString(NewLine)
		f.WriteString("Total Packet Receive : " + fmt.Sprintf("%d", totalP))
		f.WriteString(NewLine)
		f.WriteString("Packet Size : " + fmt.Sprintf("%d Byte", PSize))
		f.WriteString(NewLine)
		f.WriteString("Total Data Byte Receive: " + fmt.Sprintf("%d Bytes ", totalSize))
		f.WriteString(NewLine)
		f.WriteString("Protocol : Pure UDP")
		f.WriteString(NewLine)
		f.WriteString("Stop Receiving Time : " + time.Now().String())
		f.WriteString(NewLine)
		f.WriteString("Time Escaped To Receiving : " + fmt.Sprintf("%d", t2-t1) + " Milliseconds")
		f.WriteString(NewLine)
		f.WriteString("==============================")
		f.Close()
	}()

	f.WriteString(NewLine)
	f.WriteString(NewLine)
	f.WriteString(NewLine)

	f.WriteString("==============================")
	f.WriteString(NewLine)
	f.WriteString("Receiving Data From " + conn.LocalAddr().String() + NewLine)
	f.WriteString("Start Receiving Time : " + time.Now().String())
	f.WriteString(NewLine)

	for {
		lengthR, err := conn.Read(buf)
		msg := Utils.Deserialize(buf[:lengthR])
		i := makeTimestamp()

		totalSize += uint64(lengthR)

		if totalP == 0 {
			totalP = msg.MaxPacketCount
		}
		if PSize == 0 {
			PSize = uint64(lengthR)
		}

		DiffMap[i-msg.Milis] = DiffMap[i-msg.Milis] + 1

		if err != nil {
			log.Fatal(err)
		}
		if msg.SequenceNumber == msg.MaxPacketCount-1 {
			break
		}
	}
}

func MultiplexKCP(conn *kcp.UDPSession) {

	t1 := makeTimestamp()
	_ = os.Mkdir("Logs", os.ModePerm)
	f, err := os.OpenFile("Logs/"+conn.LocalAddr().String()+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("File Not Opened!!")
	}

	buf := make([]byte, defaultBufferSize)
	var totalP uint64 = 0
	var totalSize uint64 = 0
	var PSize uint64 = 0

	//map [time Diff] = Count
	var DiffMap = make(map[int64]uint64)
	defer func() {
		t2 := makeTimestamp()
		f.WriteString(NewLine)

		// Sort Result
		keys := make([]int64, 0, len(DiffMap))
		for k := range DiffMap {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, k := range keys {
			txt := fmt.Sprintf("With %d Milliseconds Different -> Count is: %d \r\n", k, DiffMap[k])
			f.WriteString(txt)
		}
		log.Println("DONE")
		f.WriteString(NewLine)
		f.WriteString(NewLine)
		f.WriteString("Total Packet Receive : " + fmt.Sprintf("%d", totalP))
		f.WriteString(NewLine)
		f.WriteString("Packet Size : " + fmt.Sprintf("%d Byte", PSize))
		f.WriteString(NewLine)
		f.WriteString("Total Data Byte Receive: " + fmt.Sprintf("%d Bytes ", totalSize))
		f.WriteString(NewLine)
		f.WriteString("Protocol : KCP")
		f.WriteString(NewLine)
		f.WriteString("Stop Receiving Time : " + time.Now().String())
		f.WriteString(NewLine)
		f.WriteString("Time Escaped To Receiving : " + fmt.Sprintf("%d", t2-t1) + " Milliseconds")
		f.WriteString(NewLine)
		f.WriteString("==============================")
		f.Close()
	}()

	f.WriteString(NewLine)
	f.WriteString(NewLine)
	f.WriteString(NewLine)

	f.WriteString("==============================")
	f.WriteString(NewLine)
	f.WriteString("Receiving Data From " + conn.LocalAddr().String() + NewLine)
	f.WriteString("Start Receiving Time : " + time.Now().String())
	f.WriteString(NewLine)

	for {
		lengthR, err := conn.Read(buf)
		msg := Utils.Deserialize(buf[:lengthR])
		i := makeTimestamp()

		totalSize += uint64(lengthR)

		if totalP == 0 {
			totalP = msg.MaxPacketCount
		}
		if PSize == 0 {
			PSize = uint64(lengthR)
		}

		DiffMap[i-msg.Milis] = DiffMap[i-msg.Milis] + 1

		if err != nil {
			log.Fatal(err)
		}
		if msg.SequenceNumber == msg.MaxPacketCount-1 {
			break
		}
	}
}

func MultiplexQUIC(session quic.Session) {

	stream, err := session.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}
	t1 := makeTimestamp()
	_ = os.Mkdir("Logs", os.ModePerm)
	f, err := os.OpenFile("Logs/"+session.LocalAddr().String()+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("File Not Opened!!")
	}

	buf := make([]byte, defaultBufferSize)
	var totalP uint64 = 0
	var totalSize uint64 = 0

	//map [time Diff] = Count
	var DiffMap = make(map[int64]uint64)
	defer func() {
		t2 := makeTimestamp()
		f.WriteString(NewLine)
		// Sort Result
		keys := make([]int64, 0, len(DiffMap))
		for k := range DiffMap {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
		for _, k := range keys {
			txt := fmt.Sprintf("With %d Milliseconds Different -> Count is: %d \r\n", k, DiffMap[k])
			f.WriteString(txt)
		}
		log.Println("DONE")
		f.WriteString(NewLine)
		f.WriteString(NewLine)
		f.WriteString("Total Packet Receive : " + fmt.Sprintf("%d", totalP))
		f.WriteString(NewLine)
		f.WriteString("Total Data Byte Receive: " + fmt.Sprintf("%d Bytes ", totalSize))
		f.WriteString(NewLine)
		f.WriteString("Protocol : QUIC")
		f.WriteString(NewLine)
		f.WriteString("Stop Receiving Time : " + time.Now().String())
		f.WriteString(NewLine)
		f.WriteString("Time Escaped To Receiving : " + fmt.Sprintf("%d", t2-t1) + " Milliseconds")
		f.WriteString(NewLine)
		f.WriteString("==============================")
		f.Close()
	}()

	f.WriteString(NewLine)
	f.WriteString(NewLine)
	f.WriteString(NewLine)

	f.WriteString("==============================")
	f.WriteString(NewLine)
	f.WriteString("Receiving Data From " + session.LocalAddr().String() + NewLine)
	f.WriteString("Start Receiving Time : " + time.Now().String())
	f.WriteString(NewLine)

	for {
		lengthR, err := stream.Read(buf)
		msg := Utils.Deserialize(buf[:lengthR])
		i := makeTimestamp()

		totalSize += uint64(lengthR)

		if totalP == 0 {
			totalP = msg.MaxPacketCount
		}

		DiffMap[i-msg.Milis] = DiffMap[i-msg.Milis] + 1

		if err != nil {
			log.Fatal(err)
		}
		if msg.SequenceNumber == msg.MaxPacketCount-1 {
			break
		}
	}

}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
