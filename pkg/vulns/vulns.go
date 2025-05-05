package vulns

import (
	"github.com/google/osv-scanner/v2/internal/imodels"
	"github.com/google/osv-scanner/v2/internal/utility/vulns"
	"github.com/ossf/osv-schema/bindings/go/osvschema"
)

type PackageInfo = imodels.PackageInfo

func IsAffected(v osvschema.Vulnerability, pkg PackageInfo) bool {
	return vulns.IsAffected(v, pkg)
}
