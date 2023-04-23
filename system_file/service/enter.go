package service

import (
	"letgo_repo/system_file/service/common"
	"letgo_repo/system_file/service/manifest"
	"letgo_repo/system_file/service/old_service"
	"letgo_repo/system_file/service/todo"
)

type groupV1 struct {
	old_service.CodeServiceI
	manifest.ServiceManifest
	todo.ServiceTodo
	CommonService common.Service
}

var SGroupV1 = groupV1{}

func init() {
	SGroupV1.CommonService = &common.ServiceImpl{}
	SGroupV1.CodeServiceI = &old_service.CodeServiceImpl{}
	SGroupV1.ServiceManifest = &manifest.ServiceManifestImpl{}
	SGroupV1.ServiceTodo = &todo.ServiceTodoImpl{}
}
