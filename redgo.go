package redgo

import (
	"bufio"
	"bytes"
	"errors"
	"net"
	"strconv"
	"time"
	"fmt"
)

type Conn struct {
	conn   net.Conn
	reader *bufio.Reader
}

const (
	Sep         = "\r\n"
	MaxArgSize  = 64000
	MaxArgCount = 64
)

var (
	errFormat = errors.New("format error")
	errSize   = errors.New("size out of bounds")
)

type Status string
type Error string
type Integer int

func Dial(addr string) (*Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, time.Second*5)
	if err != nil {
		return nil, err
	}
	return &Conn{conn: conn , reader: bufio.NewReader(conn)}, nil
}

func (conn *Conn) Close() error {
	return conn.conn.Close()
}

func (conn *Conn) Excute(args ...string) error {
	var buf bytes.Buffer
	buf.WriteString("*")
	buf.WriteString(strconv.Itoa(len(args)))
	buf.WriteString(Sep)
	for _, arg := range args {
		buf.WriteString("$")
		buf.WriteString(strconv.Itoa(len(arg)))
		buf.WriteString(Sep)
		buf.WriteString(arg)
		buf.WriteString(Sep)
	}
	_, err := conn.conn.Write(buf.Bytes())
	// fmt.Println(buf.Bytes())
	return err
}

func (conn *Conn) ReadRaw() (string, error) {
	b := make([]byte, 4096, 4096)
	n, err := conn.conn.Read(b)
	// fmt.Println(n)
	return string(b[0:n]), err
}

//func (conn *Conn) Response()(res interface{}, error){
//	b, err := conn.reader.ReadByte()
//	if err != nil {
//		return nil, err
//	}

//	switch b {
//		// OK
//		case '+':
//		line, isPrefix, err := conn.reader.ReadLine()
//		if err != nil {
//			return nil, err
//		}
//		if isPrefix{
//			return nil, errSize
//		}
//		return Status(line), nil
//		// Error
//		case '-':
//		line, isPrefix, err := c.r.ReadLine()
//		if err != nil {
//			return nil, err
//		}
//		if isPrefix {
//			return nil, ErrSize
//		}
//		return Error(line), nil
//		//Integer
//		case ':':
//		line, isPrefix, err := c.r.ReadLine()
//		if err != nil {
//			return nil, err
//		}
//		if isPrefix {
//			return nil, ErrSize
//		}
//		i, err := strconv.Atoi(string(line))
//		if err != nil {
//			return nil, ErrFormat
//		}
//		return Integer(i), nil

//		case '$':
//		conn.reader.UnreadByte()
//		return conn.ReadString()

//		case '*':
//		conn.reader.UnreadByte()
//		return conn.ReadStrings()
//	}
//}

//func (conn *Conn)ReadString()(string, error){

//}
