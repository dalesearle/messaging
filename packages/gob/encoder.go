package gob

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"playground/messaging"
	"time"
)

type encoder struct {
	pkg messaging.Package
}

func NewEncoder(pkg messaging.Package) *encoder {
	gob.Register(time.Time{})
	return &encoder{
		pkg: pkg,
	}
}

func (e *encoder) EncodePackage() error {
	var buf = new(bytes.Buffer)
	var encodedContent []byte
	var err error

	if encodedContent, err = e.pkg.EncodedContent(); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.BigEndian, uint16(e.pkg.ContentID())); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.BigEndian, e.pkg.PackageType()); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.BigEndian, e.pkg.Postmark()); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.BigEndian, e.pkg.ReturnAddress()); err != nil {
		return err
	}
	if err = binary.Write(buf, binary.BigEndian, e.pkg.Vertical()); err != nil {
		return err
	}
	if _, err = buf.Write(encodedContent); err != nil {
		return err
	}
	e.pkg.SetEncodedPackage(buf.Bytes())
	return nil
}

func (e *encoder) EncodeContent() error {
	var buf = new(bytes.Buffer)
	var content messaging.Content
	var err error

	if content, err = e.pkg.Content(); err != nil {
		return err
	}
	if err = gob.NewEncoder(buf).Encode(content); err != nil {
		return err
	}
	e.pkg.SetEncodedContent(buf.Bytes())
	return nil
}
