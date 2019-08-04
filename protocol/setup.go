package protocol

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"github.com/lucas-clemente/quic-go"
	"math/big"
	"os"
)

//func Start() {
//	udpAddr, err := net.ResolveUDPAddr("udp", endpoint)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	conn, err := net.ListenUDP("udp", udpAddr)
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Println("UDP Server:\tSUCCESS")
//	defer conn.Close()
//
//	for{
//	 	go multiplexer.Multiplex(conn)
//	}
//}

//func StartKCP(){
//	key := pbkdf2.Key([]byte("demo pass"), []byte("demo salt"), 1024, 32, sha1.New)
//	block, _ := kcp.NewAESBlockCrypt(key)
//	if listener, err := kcp.ListenWithOptions(os.Getenv("ENDPOINT"), block, 10, 3); err == nil {
//		log.Println("KCP Server:\tSUCCESS")
//
//		defer listener.Close()
//		for {
//			s, err := listener.AcceptKCP()
//			log.Println("onAccept")
//			if err != nil {
//				log.Fatal(err)
//			}
//			go multiplexer.Multiplex(s)
//		}
//	} else {
//		log.Fatal(err)
//	}
//}

func StartQUIC() {
	listener, err := quic.ListenAddr(os.Getenv("ENDPOINT"), generateTLSConfig(), nil)
	if err != nil {
		//	return err
	}
	sess, err := listener.Accept(context.Background())
	if err != nil {
		//return err
	}
	_, err = sess.AcceptStream(context.Background())
	if err != nil {
		panic(err)
	}
	// Echo through the loggingWriter
	//_, err = io.Copy(loggingWriter{stream}, stream)
	//return err
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
