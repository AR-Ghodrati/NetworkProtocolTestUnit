package Client

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/lucas-clemente/quic-go"
	"gsm/Models"
	"gsm/Utils"
	"log"
	"math/big"
	"os"
	"time"
)

func RunQUIC(count uint64) {

	tlsConf := &tls.Config{
		InsecureSkipVerify: true,
		NextProtos:         []string{"quic-echo-example"},
	}
	// wait for server to become ready
	time.Sleep(time.Second)
	log.Println("Connecting....")

	// dial to the echo server
	if sess, err := quic.DialAddr(os.Getenv("ENDPOINT"), tlsConf, nil); err == nil {

		stream, err := sess.OpenStreamSync(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			log.Println("Send", count, "Packet DONE!!")
		}()
		var i uint64 = 0
		for i = 0; i < count; i++ {
			data := makeTimestamp()
			random, _ := Utils.GenerateRandomString(200)
			_, _ = stream.Write(Utils.Serialize(Models.Message{SequenceNumber: uint64(i), Milis: data, Msg: random, MaxPacketCount: count}))
			time.Sleep(time.Microsecond)
		}
	} else {
		log.Fatal(err)
	}
}
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
