package manifest

import (
	model_type_def "letgo_repo/system_file/data_access/manifest/type_def"
	"letgo_repo/system_file/service/manifest/type_def"
)

type ServiceManifest interface {
	Get(model_type_def.QueryWrapper) type_def.Manifests
	GetList() type_def.Manifests
	Save(manifest type_def.Manifest)
}
