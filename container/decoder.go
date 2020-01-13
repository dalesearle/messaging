package container

import (
	"encoding/binary"
	"errors"
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

func (d *decoder) DecodeContainer(pkg messaging.Package) error {
	var conn = d.conn
	var err error
	var pre = make([]byte, 15)
	var pkgType = make([]byte, 1)
	var pkgLen = make([]byte, 4)
	var read int

	if read, err = conn.Read(pre); err != nil {
		return err
	}
	if read != 15 || string(pre) != messaging.Preamble {
		return errors.New("failed to read preamble") //now figure out how to resync the connection
		conn.Close()
	}
	//TODO: something with read
	if read, err = conn.Read(pkgType); err != nil {
		return err
	}
	//TODO: something with read
	if read, err = conn.Read(pkgLen); err != nil {
		return err
	}
	encLen := binary.BigEndian.Uint32(pkgLen)
	enc := make([]byte, encLen)
	//TODO: something with read
	if read, err = conn.Read(enc); err != nil {
		return err
	}
	pkg.SetPackageType(messaging.PackageType(pkgType[0])).
		SetEncodedPackage(enc)
	if err = packages.GetPackageDecoder(pkg).DecodePackage(); err != nil {
		return err
	}

	return nil
}
