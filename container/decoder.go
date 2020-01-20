package container

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"playground/messaging"
	"playground/messaging/packages"
)

type decoder struct {
	conn net.Conn
}

func NewDecoder(conn net.Conn) *decoder {
	return &decoder{
		conn: conn,
	}
}

func (d *decoder) DecodeContainer(pkg messaging.Package) (int, error) {
	var conn = d.conn
	var err error
	var pre = make([]byte, 15)
	var pkgType = make([]byte, 1)
	var pkgLen = make([]byte, 4)
	var read int
	var totalRead int

	if read, err = conn.Read(pre); err != nil {
		return 0, err
	}
	totalRead += read
	if read != 15 || string(pre) != messaging.Preamble {
		return 0, errors.New("failed to read preamble") //now figure out how to resync the connection
		conn.Close()
	}
	if read, err = conn.Read(pkgType); err != nil {
		return totalRead, err
	}
	totalRead += read
	if read, err = conn.Read(pkgLen); err != nil {
		return totalRead, err
	}
	totalRead += read
	encLen := binary.BigEndian.Uint32(pkgLen)
	fmt.Printf("dddddddlen: %d\n", uint32(encLen))
	enc := make([]byte, encLen)
	if read, err = conn.Read(enc); err != nil {
		return totalRead, err
	}
	totalRead += read
	pkg.SetPackageType(messaging.PackageType(pkgType[0])).
		SetEncodedPackage(enc)
	if err = packages.GetPackageDecoder(pkg).DecodePackage(); err != nil {
		return totalRead, err
	}

	return totalRead, nil
}
