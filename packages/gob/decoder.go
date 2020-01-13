package gob

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"playground/messaging"
	content2 "playground/messaging/content"
)

type decoder struct {
	pkg messaging.Package
}

func NewDecoder(pkg messaging.Package) *decoder {
	return &decoder{
		pkg: pkg,
	}
}

func (d *decoder) DecodePackage() error {
	var err error
	var contentID uint16
	var encodedPkg []byte
	var packageType byte
	var postmark int64
	var returnAddress uint32
	var vertical byte
	var pkg = d.pkg

	if encodedPkg, err = pkg.EncodedPackage(); err != nil {
		return err
	}
	buf := bytes.NewBuffer(encodedPkg)
	if err = binary.Read(buf, binary.BigEndian, contentID); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, packageType); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, postmark); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, returnAddress); err != nil {
		return err
	}
	if err = binary.Read(buf, binary.BigEndian, vertical); err != nil {
		return err
	}
	pkg.SetContentID(messaging.ContentID(contentID)).
		SetEncodedContent(buf.Bytes()).
		SetPackageType(messaging.PackageType(packageType)).
		SetPostmark(postmark).
		SetReturnAddress(returnAddress).
		SetVertical(messaging.Vertical(vertical))
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
	content := content2.GetContent(d.pkg.ContentID())
	if err := enc.Decode(content); err != nil {
		return nil, err
	}
	return content, nil
}
