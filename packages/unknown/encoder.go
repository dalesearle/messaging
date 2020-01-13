package unknown

import (
	"errors"
	"fmt"
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

func (e *encoder) EncodePackage() error {
	return errors.New(fmt.Sprintf("unknown package type %v, encode failed", e.pkg.PackageType()))
}

func (e *encoder) EncodeContent() error {
	return errors.New(fmt.Sprintf("unknown package type %v, encode failed", e.pkg.PackageType()))
}
