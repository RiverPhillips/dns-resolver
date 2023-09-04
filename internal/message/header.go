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

func ParseHeader(b *bytes.Reader) (*Header, error) {
	h := new(Header)

	id, err := ReadU16(b)
	if err != nil {
		return nil, err
	}

	flags, err := ReadU16(b)
	if err != nil {
		return nil, err
	}

	qs, err := ReadU16(b)
	if err != nil {
		return nil, err
	}

	as, err := ReadU16(b)
	if err != nil {
		return nil, err
	}

	aus, err := ReadU16(b)
	if err != nil {
		return nil, err
	}

	adds, err := ReadU16(b)
	if err != nil {
		return nil, err
	}

	h.ID = id
	h.Flags = flags
	h.NumQuestions = qs
	h.NumAnswers = as
	h.NumAuthorities = aus
	h.NumAdditionals = adds

	return h, nil
}

func u16ToBytes(i uint16) []byte {
	b := make([]byte, 2)

	binary.BigEndian.PutUint16(b, i)

	return b
}

func ReadU16(b *bytes.Reader) (uint16, error) {
	uint16Bytes := make([]byte, 2)
	if _, err := b.Read(uint16Bytes); err != nil {
		return 0, err
	}

	return binary.BigEndian.Uint16(uint16Bytes), nil
}
