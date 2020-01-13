package packages

import (
	"playground/messaging"
	"playground/messaging/packages/gob"
	"playground/messaging/packages/unknown"
)

func GetPackageDecoder(pkg messaging.Package) messaging.PackageDecoder {
	switch pkg.PackageType() {
	case messaging.GobPackage:
		return gob.NewDecoder(pkg)
	default:
		return unknown.NewDecoder(pkg)
	}
}
