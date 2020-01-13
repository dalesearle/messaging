package unknown

import "playground/messaging"

type unknown struct{}

func New() *unknown {
	return new(unknown)
}

func (u *unknown) ContentID() messaging.ContentID {
	return messaging.Unknown
}

func (u *unknown) PackageType() messaging.PackageType {
	return messaging.UnknownPackage
}
