package comparer

import (
	"github.com/google/osv-scanner/internal/semantic"
	"github.com/google/osv-scanner/pkg/models"
)

type Version = semantic.Version

func Parse(str string, ecosystem models.Ecosystem) (Version, error) {
	return semantic.Parse(str, ecosystem)
}
