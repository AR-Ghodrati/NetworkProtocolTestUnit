package utils

import (
	"../models"
	"github.com/vmihailenco/msgpack"
)

func Serialize(input models.Message) []byte {
	output, err := msgpack.Marshal(input)
	if err != nil {
		panic(err)
	}

	return output
}

func Deserialize(input []byte) models.Message {
	var msg models.Message
	err := msgpack.Unmarshal(input, &msg)
	if err != nil {
		panic(err)
	}
	return msg
}
