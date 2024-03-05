//go:build !linux
// +build !linux

package kcp

func (s *UDPSession) tx(txqueue []Message) {
	s.defaultTx(txqueue)
}
