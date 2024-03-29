package data_access

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"letgo_repo/system_file/data_access/models"
)

var MysqlDB *gorm.DB

func init() {
	//dsn := "root:123456@tcp(ali2-1.s:3306)/letgo?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "user_letgo:letgo_XAmiaoYes123_@tcp(192.168.177.130:3306)/letgo?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "user_letgo:letgo_XAmiaoYes123_@tcp(ali2-1.s:3306)/letgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	MysqlDB = db
}

func GenerateTable() error {
	migrator := MysqlDB.Migrator()
	err := migrator.AutoMigrate(&models.TopicTags{})
	if err != nil {
		return err
	}

	err = migrator.AutoMigrate(&models.TodoQuestion{})
	if err != nil {
		return err
	}

	err = migrator.AutoMigrate(&models.Todo{})
	if err != nil {
		return err
	}

	err = migrator.AutoMigrate(&models.Manifest{})
	if err != nil {
		return err
	}

	err = migrator.AutoMigrate(&models.QuestionTest{})
	if err != nil {
		return err
	}

	err = migrator.AutoMigrate(&models.ToDoQuestion{})
	if err != nil {
		return err
	}

	err = migrator.AutoMigrate(&models.OperationRecords{})
	if err != nil {
		return err
	}

	err = migrator.AutoMigrate(&models.TopCompanyTags{})
	if err != nil {
		return err
	}

	err = migrator.AutoMigrate(&models.Question{})
	if err != nil {
		return err
	}

	return nil
}
