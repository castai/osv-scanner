package models

import (
	"github.com/google/osv-scalibr/extractor"
	"github.com/google/osv-scanner/v2/internal/imodels"
)

type ModelsPackageInfo = imodels.PackageInfo

func FromInventory(inventory *extractor.Package) ModelsPackageInfo {
	return imodels.FromInventory(inventory)
}
