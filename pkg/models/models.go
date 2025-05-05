package models

import (
	"github.com/google/osv-scalibr/extractor"
	"github.com/google/osv-scanner/v2/internal/imodels"
)

func FromInventory(inventory *extractor.Package) PackageInfo {
	return imodels.FromInventory(inventory)
}
