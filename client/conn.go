package client

import (
	"net"
	"time"
)

// Addr 等价net.Addr
type Addr = net.Addr

// Conn socks5连接
type Conn struct {
	conn net.Conn
}

// Read 从proxy中读取数据
func (c *Conn) Read(b []byte) (n int, err error) {
	return c.conn.Read(b)
}

// Write 往proxy中写数据
func (c *Conn) Write(b []byte) (n int, err error) {
	return c.conn.Write(b)
}

// Close 关闭proxy连接
func (c *Conn) Close() error {
	return c.conn.Close()
}

// LocalAddr 返回本地地址
func (c *Conn) LocalAddr() Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr 返回远程地址
func (c *Conn) RemoteAddr() Addr {
	return c.conn.RemoteAddr()
}

// SetDeadline 设置读写超时时间，零值为不限
func (c *Conn) SetDeadline(t time.Time) error {
	return c.conn.SetDeadline(t)
}

// SetReadDeadline 设定该连接的读操作deadline，参数t为零值表示不设置期限
func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.conn.SetReadDeadline(t)
}

// SetWriteDeadline 设定该连接的写操作deadline，参数t为零值表示不设置期限
func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.conn.SetWriteDeadline(t)
}
