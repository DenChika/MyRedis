package lib

import (
	"net"
	"sync"
)

type SafeConn struct {
	Conn net.Conn
	Rwmu *sync.RWMutex
}

func NewSafeConn(Conn net.Conn) *SafeConn {
	return &SafeConn{Conn: Conn, Rwmu: &sync.RWMutex{}}
}

func (c *SafeConn) Read(b []byte) (n int, err error) {
	return c.Conn.Read(b)
}

func (c *SafeConn) Write(b []byte) (n int, err error) {
	c.Rwmu.Lock()
	defer c.Rwmu.Unlock()
	return c.Conn.Write(b)
}

func (c *SafeConn) Close() error {
	return c.Conn.Close()
}
