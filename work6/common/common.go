package common

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

var HeaderFlag = [2]byte{0xaa, 0xbb}

type GoimPkgHeader struct {
	HeaderFlag      [2]byte
	PackageLength   uint32
	HeaderLength    uint16
	ProtocolVersion uint16
	Operation       uint32
	SequenceId      uint32
}

const (
	HeaderLength = uint16(18)
)
const (
	SOCKET_ERROR_CLIENT_CLOSED = "Client has been closed"
	SOCKET_ERROR_SERVER_CLOSED = "Server has been closed"
	SOCKET_ERROR_TIMEOUT       = "Timeout"
)

type SocketUtil struct {
	Conn net.Conn
}

func (fd *SocketUtil) Init(conn net.Conn) {
	fd.Conn = conn
}
func (fd *SocketUtil) WritePkg(h GoimPkgHeader, data []byte) (int, error) {
	if fd == nil {
		return -1, errors.New(SOCKET_ERROR_SERVER_CLOSED)
	}
	if len(data) == 0 {
		return 0, nil
	}
	buff := bytes.NewBuffer([]byte{})
	binary.Write(buff, binary.BigEndian, []byte{0xaa, 0xbb})
	binary.Write(buff, binary.BigEndian, uint32(len(data))+uint32(HeaderLength))
	binary.Write(buff, binary.BigEndian, HeaderLength)
	binary.Write(buff, binary.BigEndian, h.ProtocolVersion)
	binary.Write(buff, binary.BigEndian, h.Operation)
	binary.Write(buff, binary.BigEndian, h.SequenceId)

	binary.Write(buff, binary.BigEndian, data)
	allBytes := buff.Bytes()
	return fd.writeNByte(allBytes)
}
func (fd *SocketUtil) ReadPkg() ([]byte, error) {
	if fd == nil || fd.Conn == nil {
		return nil, errors.New(SOCKET_ERROR_SERVER_CLOSED)
	}
	head, err := fd.readHead()
	if err != nil {
		return nil, err
	}
	if head.HeaderFlag != [2]byte{0xaa, 0xbb} {
		return nil, errors.New("Head package error")
	}
	datas, err := fd.readNByte(head.PackageLength - uint32(head.HeaderLength))
	if err != nil {
		return nil, err
	}
	return datas, nil
}
func (fd *SocketUtil) readHead() (*GoimPkgHeader, error) {
	data, err := fd.readNByte(uint32(HeaderLength))
	if err != nil {
		return nil, err
	}
	h := GoimPkgHeader{}
	buff := bytes.NewBuffer(data)
	binary.Read(buff, binary.BigEndian, &h.HeaderFlag)
	binary.Read(buff, binary.BigEndian, &h.PackageLength)
	binary.Read(buff, binary.BigEndian, &h.HeaderLength)
	binary.Read(buff, binary.BigEndian, &h.ProtocolVersion)
	binary.Read(buff, binary.BigEndian, &h.Operation)
	binary.Read(buff, binary.BigEndian, &h.SequenceId)
	return &h, nil
}
func (fd *SocketUtil) readNByte(n uint32) ([]byte, error) {
	data := make([]byte, n)
	for x := 0; x < int(n); {
		length, err := fd.Conn.Read(data[x:])
		if length == 0 {
			return nil, errors.New(SOCKET_ERROR_CLIENT_CLOSED)
		}
		if err != nil {
			return nil, err
		}
		x += length
	}
	return data, nil
}
func (fd *SocketUtil) writeNByte(data []byte) (int, error) {
	n, err := fd.Conn.Write(data)
	if err != nil {
		return -1, err
	} else {
		return n, nil
	}
}
func (fd *SocketUtil) Close() {
	fd.Conn.Close()
}
