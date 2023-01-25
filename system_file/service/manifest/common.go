package manifest

import "letgo_repo/system_file/service/manifest/type_def"

type ServiceManifest interface {
	GetList() type_def.Manifests
	Save(manifest type_def.Manifest)
}
