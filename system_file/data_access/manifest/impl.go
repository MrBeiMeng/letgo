package manifest

import (
	"gorm.io/gorm/clause"
	"letgo_repo/system_file/data_access"
	"letgo_repo/system_file/data_access/manifest/type_def"
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

func (i DAManifestImpl) Select(queryWrapper type_def.QueryWrapper) (resultList []models.Manifest) {
	if queryWrapper.CaseNothing() {
		// 查询所有
		err := data_access.MysqlDB.Order("tags").Find(&resultList).Error
		if err != nil {
			panic(err)
		}

		return
	}

	tmpResult1 := make([]models.Manifest, 0)
	if len(queryWrapper.TagSlice) != 0 {
		err := data_access.MysqlDB.Where("tags in ?", queryWrapper.TagSlice).Find(&tmpResult1).Error
		if err != nil {
			panic(err)
		}
	}

	tmpResult2 := make([]models.Manifest, 0)
	if len(queryWrapper.TitleSlice) != 0 {
		err := data_access.MysqlDB.Where("title in ?", queryWrapper.TitleSlice).Find(&tmpResult2).Error
		if err != nil {
			panic(err)
		}
	}

	for _, manifest := range tmpResult1 {
		resultList = append(resultList, manifest)
	}
	for _, manifest := range tmpResult2 {
		resultList = append(resultList, manifest)
	}
	return
}

func (DAManifestImpl) test() {
	//TODO implement me
	panic("implement me")
}
