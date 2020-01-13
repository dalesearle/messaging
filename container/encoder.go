package container

import (
	"bytes"
	"encoding/binary"
	"playground/messaging"
)

type encoder struct {
	pkg messaging.Package
}

func NewEncoder(pkg messaging.Package) *encoder {
	return &encoder{
		pkg: pkg,
	}
}

func (e *encoder) EncodeContainer(buf *bytes.Buffer) error {
	var encodedPackage []byte
	var encodedLen = make([]byte, 4)
	var err error

	if encodedPackage, err = e.pkg.EncodedPackage(); err != nil {
		return err
	}
	if _, err = buf.WriteString(messaging.Preamble); err != nil {
		return err
	}
	if err = buf.WriteByte(byte(e.pkg.PackageType())); err != nil {
		return err
	}
	binary.BigEndian.PutUint32(encodedLen, uint32(len(encodedPackage)))
	if _, err = buf.Write(encodedLen); err != nil {
		return err
	}
	if _, err = buf.Write(encodedPackage); err != nil {
		return err
	}
	return nil
}
