package kcp

import "net"

type Message struct {
	// When writing, the Buffers field must contain at least one
	// byte to write.
	// When reading, the Buffers field will always contain a byte
	// to read.
	Buffers [][]byte

	// OOB contains protocol-specific control or miscellaneous
	// ancillary data known as out-of-band data.
	OOB []byte

	// Addr specifies a destination address when writing.
	// It can be nil when the underlying protocol of the raw
	// connection uses connection-oriented communication.
	// After a successful read, it may contain the source address
	// on the received packet.
	Addr net.Addr

	N     int // # of bytes read or written from/to Buffers
	NN    int // # of bytes read or written from/to OOB
	Flags int // protocol-specific information on the received message
}
