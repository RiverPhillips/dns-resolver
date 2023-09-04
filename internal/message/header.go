package message

import (
	"bytes"
	"encoding/binary"
)

type Header struct {
	ID             uint16
	Flags          uint16
	NumQuestions   uint16
	NumAnswers     uint16
	NumAuthorities uint16
	NumAdditionals uint16
}

func (h *Header) ToBytes() []byte {
	buf := new(bytes.Buffer)

	buf.Write(u16ToBytes(h.ID))
	buf.Write(u16ToBytes(h.Flags))
	buf.Write(u16ToBytes(h.NumQuestions))
	buf.Write(u16ToBytes(h.NumAnswers))
	buf.Write(u16ToBytes(h.NumAuthorities))
	buf.Write(u16ToBytes(h.NumAdditionals))

	return buf.Bytes()
}

func u16ToBytes(i uint16) []byte {
	b := make([]byte, 2)

	binary.BigEndian.PutUint16(b, i)

	return b
}
