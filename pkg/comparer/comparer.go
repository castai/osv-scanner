package comparer

import (
	"github.com/google/osv-scanner/internal/semantic"
	"github.com/google/osv-scanner/pkg/models"
	"github.com/google/osv-scanner/v2/internal/imodels"
	"github.com/google/osv-scanner/v2/internal/utility/vulns"
	"github.com/ossf/osv-schema/bindings/go/osvschema"
)

type Version = semantic.Version

func Parse(str string, ecosystem models.Ecosystem) (Version, error) {
	return semantic.Parse(str, ecosystem)
}

func IsAffected(v osvschema.Vulnerability, pkg imodels.PackageInfo) bool {
	return vulns.IsAffected(v, pkg)
}
