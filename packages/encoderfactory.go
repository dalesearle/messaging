package packages

import (
	"playground/messaging"
	"playground/messaging/packages/gob"
	"playground/messaging/packages/unknown"
)

func GetPackageEncoder(pkg messaging.Package) messaging.PackageEncoder {
	switch pkg.PackageType() {
	case messaging.GobPackage:
		return gob.NewEncoder(pkg)
	default:
		return unknown.NewEncoder(pkg)
	}
}
