package net

import (
	"fmt"
	"sync"
	"time"
)

type TCP struct {
	once sync.Once
	conn *conn

	host string
}

func NewTCP(host string) *TCP {
	return &TCP{host: host}
}

func (tcp *TCP) GetConnection() *conn {
	tcp.once.Do(func() {
		tcp.conn = connect(tcp.host)
	})

	return tcp.conn
}

type conn struct {
	host        string
	connectedAt time.Time
}

func connect(host string) *conn {
	// Дуже важна логіка для реального конекту :)
	return &conn{host: host, connectedAt: time.Now()}
}

func (c conn) Print() {
	fmt.Println(c.connectedAt.Format(time.Stamp), "-час коли ми ступили на", c.host)
}
