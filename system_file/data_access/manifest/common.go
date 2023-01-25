package manifest

import (
	"letgo_repo/system_file/data_access/manifest/type_def"
	"letgo_repo/system_file/data_access/models"
)

type DAManifest interface {
	test()
	Select(type_def.QueryWrapper) []models.Manifest
	InsertManifest(modelManifest models.Manifest)
}

var DAManifestV1 DAManifest = &DAManifestImpl{}
