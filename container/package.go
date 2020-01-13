package container

import (
	"playground/messaging"
	"playground/messaging/packages"
	"time"
)

var _ messaging.Package = new(Package)

type Package struct {
	content        messaging.Content
	contentID      messaging.ContentID
	encodedContent []byte
	encodedPackage []byte
	packageType    messaging.PackageType
	postmark       int64
	returnAddress  uint32
	vertical       messaging.Vertical
}

func NewPackage() messaging.Package {
	return &Package{
		postmark: time.Now().In(time.UTC).Unix(),
	}
}

func (b *Package) SetContent(content messaging.Content) messaging.Package {
	b.content = content
	b.SetContentID(content.ContentID())
	b.SetPackageType(content.PackageType())
	return b
}

func (b *Package) SetContentID(id messaging.ContentID) messaging.Package {
	b.contentID = id
	return b
}

func (b *Package) SetEncodedContent(bytes []byte) messaging.Package {
	b.encodedContent = bytes
	return b
}

func (b *Package) SetEncodedPackage(bytes []byte) messaging.Package {
	b.encodedPackage = bytes
	return b
}

func (b *Package) SetPackageType(t messaging.PackageType) messaging.Package {
	b.packageType = t
	return b
}

func (b *Package) SetPostmark(p int64) messaging.Package {
	b.postmark = p
	return b
}

func (b *Package) SetReturnAddress(addr uint32) messaging.Package {
	b.returnAddress = addr
	return b
}

func (b *Package) SetVertical(v messaging.Vertical) messaging.Package {
	b.vertical = v
	return b
}

func (b *Package) Content() (messaging.Content, error) {
	var err error
	var content messaging.Content

	if b.content == nil {
		if content, err = packages.GetPackageDecoder(b).DecodeContent(); err != nil {
			return nil, err
		}
		b.content = content
	}
	return b.content, nil
}

func (b *Package) ContentID() messaging.ContentID {
	return b.contentID
}

func (b *Package) EncodedContent() ([]byte, error) {
	if b.encodedContent == nil {
		if err := packages.GetPackageEncoder(b).EncodeContent(); err != nil {
			return nil, err
		}
	}
	return b.encodedContent, nil
}

func (b *Package) EncodedPackage() ([]byte, error) {
	if b.encodedPackage == nil {
		if err := packages.GetPackageEncoder(b).EncodePackage(); err != nil {
			return nil, err
		}
	}
	return b.encodedPackage, nil
}

func (b *Package) PackageType() messaging.PackageType {
	return b.packageType
}

func (b *Package) ReturnAddress() uint32 {
	return b.returnAddress
}

func (b *Package) Postmark() int64 {
	return b.postmark
}

func (b *Package) Vertical() messaging.Vertical {
	return b.vertical
}

//
