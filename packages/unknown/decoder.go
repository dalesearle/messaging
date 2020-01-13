package unknown

import (
	"errors"
	"fmt"
	"playground/messaging"
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
	return errors.New(fmt.Sprintf("unknown package type %v, decode failed", d.pkg.PackageType()))
}

func (d *decoder) DecodeContent() (messaging.Content, error) {
	return nil, errors.New(fmt.Sprintf("unknown package type %v, decode failed", d.pkg.PackageType()))
}
