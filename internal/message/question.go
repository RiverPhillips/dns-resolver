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

func ParseQuestion(b *bytes.Reader) (*Question, error) {
	q := new(Question)

	name, err := decodeDnsName(b)
	if err != nil {
		return nil, err
	}
	class, err := ReadU16(b)
	if err != nil {
		return nil, err
	}

	type_, err := ReadU16(b)
	if err != nil {
		return nil, err
	}

	q.name = name
	q.class = class
	q.type_ = type_

	return q, nil
}

func decodeDnsName(b *bytes.Reader) ([]byte, error) {
	decoded := []byte{}

	for {
		segmentLen, err := b.ReadByte()
		if err != nil {
			return nil, err
		}
		if segmentLen == 0 {
			break
		}

		segment := make([]byte, segmentLen)
		if _, err := b.Read(segment); err != nil {
			return nil, err
		}
		segment = append(segment, '.')
		decoded = append(decoded, segment...)
	}
	return decoded, nil
}
