package message

import "math/rand"

const (
	RecursionDesired = 1 << 8
	IN               = 1
	TYPE_A           = 1
)

func BuildQuery(name string, recordType uint16) []byte {
	id := generateId()
	header := Header{
		ID:           id,
		NumQuestions: 1,
		Flags:        RecursionDesired,
	}

	question := NewQuestion(name, recordType, IN)

	return append(header.ToBytes(), question.ToBytes()...)
}

func generateId() uint16 {
	return uint16(rand.Int31n(1 << 16))
}
