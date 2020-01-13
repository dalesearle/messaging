package messaging

const (
	Preamble string = "~~~~~WEAVE~~~~~"

	// PackageType
	GobPackage     PackageType = 'G'
	UnknownPackage PackageType = 'U'

	// Verticals
	ClassicVertical Vertical = 'C'

	// Consumable ID's
	Unknown   ContentID = 0
	TableData ContentID = 1
)
