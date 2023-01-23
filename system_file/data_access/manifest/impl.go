package manifest

import (
	"gorm.io/gorm/clause"
	"letgo_repo/system_file/data_access"
	"letgo_repo/system_file/data_access/models"
	utils2 "letgo_repo/system_file/utils"
	"strings"
)

type DAManifestImpl struct{}

func (i DAManifestImpl) InsertManifest(modelManifest models.Manifest) {
	// ! replace and insert
	var tmpModel models.Manifest
	data_access.MysqlDB.Where("title = ?", modelManifest.Title).First(&tmpModel)
	if tmpModel.Tags != "" {
		tmpMap := make(map[string]struct{})

		for _, tag := range strings.Split(tmpModel.Tags, ",") {
			tmpMap[tag] = struct{}{}
		}

		for _, tag := range strings.Split(modelManifest.Tags, ",") {
			tmpMap[tag] = struct{}{}
		}

		tmpArr := make([]string, 0)
		for key, _ := range tmpMap {
			tmpArr = append(tmpArr, key)
		}

		modelManifest.Tags = strings.Join(tmpArr, ",")
	}

	err := data_access.MysqlDB.Clauses(clause.OnConflict{DoNothing: true}).Create(&modelManifest).Error
	if err != nil {
		println(utils2.GetColorRed(err.Error()))
		//return true
	}
}

func (i DAManifestImpl) Select() []models.Manifest {
	//TODO implement me
	panic("implement me")
}

func (DAManifestImpl) test() {
	//TODO implement me
	panic("implement me")
}
