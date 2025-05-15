package dlt645

import (
	"io"
	"sync"

	"github.com/tarm/serial"
)

type serialPort struct {
	serial.Config
	mu     sync.Mutex
	port   io.ReadWriteCloser
	Logger ILogger
}

func (mb *serialPort) Connect() (err error) {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	return mb.connect()
}

func (mb *serialPort) connect() error {
	if mb.port == nil {
		port, err := serial.OpenPort(&mb.Config)
		if err != nil {
			return err
		}
		mb.port = port
	}
	return nil
}

func (mb *serialPort) Close() (err error) {
	mb.mu.Lock()
	defer mb.mu.Unlock()

	return mb.close()
}

func (mb *serialPort) close() (err error) {
	if mb.port != nil {
		err = mb.port.Close()
		mb.port = nil
	}
	return
}

func (mb *serialPort) logf(format string, v ...interface{}) {
	if mb.Logger != nil {
		mb.Logger.Printf(format, v...)
	}
}
