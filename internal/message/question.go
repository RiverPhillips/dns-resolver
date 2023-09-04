package message

import (
	"bytes"
	"strings"
)

func EncodeDnsName(name string) []byte {
	encoded := []byte{}
	for _, seg := range strings.Split(name, ".") {
		len := len(seg)
		if len > 0 {
			encoded = append(encoded, byte(len))
			encoded = append(encoded, []byte(seg)...)
		}
	}
	encoded = append(encoded, 0)
	return encoded
}

type Question struct {
	name  []byte
	type_ uint16
	class uint16
}

func NewQuestion(name string, type_ uint16, class uint16) *Question {
	return &Question{
		name:  EncodeDnsName(name),
		type_: type_,
		class: class,
	}
}

func (q *Question) ToBytes() []byte {
	buf := new(bytes.Buffer)

	buf.Write(q.name)
	buf.Write(u16ToBytes(q.type_))
	buf.Write(u16ToBytes(q.class))

	return buf.Bytes()
}
