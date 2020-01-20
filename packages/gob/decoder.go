package gob

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"playground/messaging"
	"playground/messaging/content"
	"time"
)

type decoder struct {
	pkg messaging.Package
}

func NewDecoder(pkg messaging.Package) *decoder {
	gob.Register(time.Time{})
	return &decoder{
		pkg: pkg,
	}
}

func (d *decoder) DecodePackage() error {
	var err error
	var contentID messaging.ContentID
	var encodedPkg []byte
	var pkgType messaging.PackageType
	var postmark int64
	var returnAddress uint32
	var vertical messaging.Vertical
	var pkg = d.pkg

	if encodedPkg, err = pkg.EncodedPackage(); err != nil {
		return err
	}
	buf := bytes.NewBuffer(encodedPkg)
	if err = binary.Read(buf, binary.BigEndian, &contentID); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, &pkgType); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, &postmark); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, &returnAddress); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, &vertical); err != nil {
		return err
	}
	pkg.SetContentID(contentID).
		SetEncodedContent(buf.Bytes()).
		SetPackageType(pkgType).
		SetPostmark(postmark).
		SetReturnAddress(returnAddress).
		SetVertical(vertical)
	return nil
}

func (d *decoder) DecodeContent() (messaging.Content, error) {
	var encodedContent []byte
	var err error

	if encodedContent, err = d.pkg.EncodedContent(); err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(encodedContent)
	enc := gob.NewDecoder(buf)
	cont := content.GetContent(d.pkg.ContentID())
	if err := enc.Decode(cont); err != nil {
		return nil, err
	}
	return cont, nil
}
