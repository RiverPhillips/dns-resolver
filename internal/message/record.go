package message

type Record struct {
	name  []byte
	type_ uint16
	class uint16
	ttl   uint32
	data  []byte
}
