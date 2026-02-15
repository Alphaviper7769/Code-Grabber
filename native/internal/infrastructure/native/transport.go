package native

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"os"
)

type Transport struct{}

func NewTransport() *Transport {
	return &Transport{}
}

func (t *Transport) ReadMessage(v interface{}) error {
	// 1️⃣ Read 4 byte length header
	var length uint32
	err := binary.Read(os.Stdin, binary.LittleEndian, &length)
	if err != nil {
		return err
	}

	// 2️⃣ Read JSON payload
	message := make([]byte, length)
	_, err = io.ReadFull(os.Stdin, message)
	if err != nil {
		return err
	}

	// 3️⃣ Unmarshal into struct
	return json.Unmarshal(message, v)
}

func (t *Transport) WriteMessage(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	length := uint32(len(data))

	// Write 4 byte header
	err = binary.Write(os.Stdout, binary.LittleEndian, length)
	if err != nil {
		return err
	}

	// Write payload
	_, err = os.Stdout.Write(data)
	return err
}
