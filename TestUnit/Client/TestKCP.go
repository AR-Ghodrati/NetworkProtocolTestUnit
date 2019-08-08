package Client

import (
	"crypto/sha1"
	"github.com/xtaci/kcp-go"
	"golang.org/x/crypto/pbkdf2"
	"gsm/Models"
	"gsm/Utils"
	"log"
	"net"
	"os"
	"time"
)

func RunKCP(count uint64) {
	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
	block, _ := kcp.NewAESBlockCrypt(key)

	// wait for server to become ready
	time.Sleep(time.Second)
	log.Println("Connecting....")

	// dial to the echo server
	if sess, err := kcp.DialWithOptions(os.Getenv("ENDPOINT"), block, 10, 3); err == nil {
		defer func() {
			log.Println("Send", count, "Packet DONE!!")
		}()
		var i uint64 = 0
		for i = 0; i < count; i++ {
			data := makeTimestamp()
			random, _ := Utils.GenerateRandomString(200)
			_, _ = sess.Write(Utils.Serialize(Models.Message{SequenceNumber: uint64(i), Milis: data, Msg: random, MaxPacketCount: count}))
			time.Sleep(time.Microsecond)
		}
	} else {
		log.Fatal(err)
	}
}
func getIP(hostname string) string {
	addr, err := net.LookupIP(hostname)
	if err != nil {
		log.Println("Unknown host")
	} else {
		log.Println("IP address: ", addr)
		return addr[0].String()
	}
	return string("")
}
