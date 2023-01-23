package service

import "letgo_repo/system_file/service/manifest"

type GroupV1 struct {
	CodeServiceI
	manifest.ServiceManifest
}

var CodeServiceGroupV1 GroupV1 = GroupV1{}

func init() {
	CodeServiceGroupV1.CodeServiceI = &CodeServiceImpl{}
	CodeServiceGroupV1.ServiceManifest = &manifest.ServiceManifestImpl{}
}
