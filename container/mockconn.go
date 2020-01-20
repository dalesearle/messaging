package container

import (
	"bytes"
	"net"
	"time"
)

type mockConn struct {
	addr net.Addr
	rdr  *bytes.Reader
	wrtr *bytes.Buffer
}

func newMockConn(bites []byte) *mockConn {
	return &mockConn{
		addr: newMockAddr(),
		rdr:  bytes.NewReader(bites),
		wrtr: new(bytes.Buffer),
	}
}

func (c *mockConn) Read(b []byte) (n int, err error) {
	return c.rdr.Read(b)
}

func (c *mockConn) Write(b []byte) (n int, err error) {
	return c.wrtr.Write(b)
}

func (c *mockConn) Close() error {
	return nil
}

func (c *mockConn) LocalAddr() net.Addr {
	return c.addr
}

func (c *mockConn) RemoteAddr() net.Addr {
	return c.addr
}

func (c *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *mockConn) SetWriteDeadline(t time.Time) error {
	return nil
}
