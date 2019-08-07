package Utils

import (
	"github.com/vmihailenco/msgpack"
	"gsm/Models"
)

func Serialize(input Models.Message) []byte {
	output, err := msgpack.Marshal(input)
	if err != nil {
		panic(err)
	}

	return output
}

func Deserialize(input []byte) Models.Message {
	var msg Models.Message
	err := msgpack.Unmarshal(input, &msg)
	if err != nil {
		panic(err)
	}
	return msg
}
