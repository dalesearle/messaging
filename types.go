package messaging

type Content interface {
	ContentID() ContentID
	PackageType() PackageType
	//TODO: immutable content vs builder
}

type ContentID uint16

type Package interface {
	Content() (Content, error)
	ContentID() ContentID
	EncodedContent() ([]byte, error)
	EncodedPackage() ([]byte, error)
	PackageType() PackageType
	ReturnAddress() uint32
	Postmark() int64
	Vertical() Vertical
	SetContent(Content) Package
	SetContentID(ContentID) Package
	SetEncodedContent([]byte) Package
	SetEncodedPackage([]byte) Package
	SetPackageType(PackageType) Package
	SetReturnAddress(uint32) Package
	SetPostmark(int64) Package
	SetVertical(Vertical) Package
}

type PackageDecoder interface {
	DecodePackage() error
	DecodeContent() (Content, error)
}

type PackageEncoder interface {
	EncodePackage() error
	EncodeContent() error
}

type PackageType byte

type Vertical byte
