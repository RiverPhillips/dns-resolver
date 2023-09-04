package message_test

import (
	"testing"

	"github.com/RiverPhillips/dns-resolver/internal/message"
	"github.com/stretchr/testify/assert"
)

func TestEncodeDnsName(t *testing.T) {
	tc := map[string]struct {
		name string
		want []byte
	}{
		"empty": {
			name: "",
			want: []byte{0},
		},
		"google.com": {
			name: "google.com",
			want: []byte{6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0},
		},
		"www.google.com": {
			name: "www.google.com",
			want: []byte{3, 119, 119, 119, 6, 103, 111, 111, 103, 108, 101, 3, 99, 111, 109, 0},
		},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, message.EncodeDnsName(tt.name))
		})
	}
}
