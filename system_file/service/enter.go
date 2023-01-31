package service

import (
	"letgo_repo/system_file/service/manifest"
	"letgo_repo/system_file/service/old_service"
	"letgo_repo/system_file/service/todo"
)

type GroupV1 struct {
	old_service.CodeServiceI
	manifest.ServiceManifest
	todo.ServiceTodo
}

var ServiceGroupV1 GroupV1 = GroupV1{}

func init() {
	ServiceGroupV1.CodeServiceI = &old_service.CodeServiceImpl{}
	ServiceGroupV1.ServiceManifest = &manifest.ServiceManifestImpl{}
	ServiceGroupV1.ServiceTodo = &todo.ServiceTodoImpl{}
}
