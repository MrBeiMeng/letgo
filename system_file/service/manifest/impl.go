package manifest

import (
	"letgo_repo/system_file/data_access/manifest"
	"letgo_repo/system_file/service/manifest/type_def"
)

type ServiceManifestImpl struct{}

func (s ServiceManifestImpl) Save(manifestObj type_def.Manifest) {
	manifest.DAManifestV1.InsertManifest(manifestObj.ConvToModel())
}

func (s ServiceManifestImpl) GetList() {
	manifests := manifest.DAManifestV1.Select()

	resultList := make(type_def.Manifests, 0)
	resultList.ConvFromModels(manifests)

	return
}
