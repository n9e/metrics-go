package statsdlib

import (
	"fmt"
	"io"
	"os"
	"strings"
	"syscall"
	"time"
)

var _loggerFd io.Writer

type logger struct{}

func (this logger) Init() error {
	fd, err := this.mkLogFile()
	if err != nil {
		return err
	}

	_loggerFd = fd
	return nil
}

func (this logger) Info(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(_loggerFd, this.now()+" [INFO] "+format, args...)
}

func (this logger) Erro(format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(_loggerFd, this.now()+" [ERRO] "+format, args...)
}

func (this logger) mkLogFile() (io.Writer, error) {
	wd, _ := syscall.Getwd()

	logdir := wd + "/.statsd"
	os.MkdirAll(logdir, 0777)

	logfn := logdir + "/statsd.log"
	f, err := os.OpenFile(logfn, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	return f, err
}

func (this logger) now() string {
	return string(time.Now().Format("2006-01-02 15:04:05"))
}
