package imodels

import (
	"github.com/google/osv-scalibr/extractor"
	internalmodels "github.com/google/osv-scanner/v2/internal/imodels"
)

type PackageInfo = internalmodels.PackageInfo

func FromInventory(inventory *extractor.Package) PackageInfo {
	return internalmodels.FromInventory(inventory)
}
