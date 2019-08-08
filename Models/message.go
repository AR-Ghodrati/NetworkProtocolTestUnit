package Models

type Message struct {
	SequenceNumber uint64
	Msg            string
	Milis          int64
	MaxPacketCount uint64
}
