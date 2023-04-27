package manifest

import (
	"letgo_repo/system_file/data_access/manifest"
	model_type_def "letgo_repo/system_file/data_access/manifest/type_def"
	"letgo_repo/system_file/data_access/question"
	"letgo_repo/system_file/service/manifest/type_def"
)

type ServiceManifestImpl struct{}

func (s ServiceManifestImpl) Save(manifestObj type_def.Manifest) {
	manifest.DAManifestV1.InsertManifest(manifestObj.ConvToModel())
}

func (s ServiceManifestImpl) Get(queryWrapper model_type_def.QueryWrapper) type_def.Manifests {
	manifests := manifest.DAManifestV1.Select(queryWrapper)

	resultList := make(type_def.Manifests, 0)
	resultList.ConvFromModels(manifests)

	// 获取question信息
	for i := 0; i < len(resultList); i++ {
		resultList[i].Questions = question.DAQuestionV1.GetByIds(resultList[i].QuestionsFrontIds)
	}

	return resultList
}

func (s ServiceManifestImpl) GetList() type_def.Manifests {
	manifests := manifest.DAManifestV1.Select(model_type_def.QueryWrapper{})

	resultList := make(type_def.Manifests, 0)
	resultList.ConvFromModels(manifests)

	// 获取question信息
	for i := 0; i < len(resultList); i++ {
		resultList[i].Questions = question.DAQuestionV1.GetByIds(resultList[i].QuestionsFrontIds)
	}

	return resultList
}
