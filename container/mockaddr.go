package container

import "net"

type mockAddr struct{}

func newMockAddr() net.Addr {
	return new(mockAddr)
}
func (a *mockAddr) Network() string {
	return "tcp"
}

func (a *mockAddr) String() string {
	return "localhost"
}
