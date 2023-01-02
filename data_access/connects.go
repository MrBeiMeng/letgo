package data_access

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"letgo_repo/data_access/models"
)

var MysqlDB *gorm.DB

func init() {
	dsn := "root:123456@tcp(ali2-1.s:3306)/letgo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db2, _ := db.DB()
	db2.Ping()
	//db2.Close()
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

	err = migrator.AutoMigrate(&models.Questions{})
	if err != nil {
		return err
	}

	return nil
}
