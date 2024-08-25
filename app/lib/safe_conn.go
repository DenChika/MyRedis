package lib

import (
	"net"
	"sync"
)

type SafeConn struct {
	conn net.Conn
	rwmu *sync.RWMutex
}

func NewSafeConn(Conn net.Conn) *SafeConn {
	return &SafeConn{conn: Conn, rwmu: &sync.RWMutex{}}
}

func (c *SafeConn) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}

func (c *SafeConn) Write(b []byte) (n int, err error) {
	c.rwmu.Lock()
	defer c.rwmu.Unlock()
	return c.conn.Write(b)
}

func (c *SafeConn) Close() error {
	return c.conn.Close()
}
