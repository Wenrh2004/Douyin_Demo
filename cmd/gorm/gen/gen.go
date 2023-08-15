package main

import (
	"Douyin_Demo/config"
	"Douyin_Demo/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Querier interface {
	// GetByRoles query data by roles and return it as *slice of pointer*
	//   (The below blank line is required to comment for the generated method)
	//	//
	// SELECT * FROM @@table WHERE role IN @rolesName
	GetByRoles(rolesName ...string) ([]*gen.T, error)
}

func generateModelByGen() {
	gormgen := gen.NewGenerator(gen.Config{
		OutPath: "repo",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
	})

	gormDB, err := gorm.Open(mysql.Open(config.AppConfig.DSN))
	fmt.Println("gormDB", config.AppConfig.DSN)
	if err != nil {
		panic("mysql - database connect error  == > " + err.Error())
	}
	gormgen.UseDB(gormDB)

	// generate basic DAO API for struct 'model.Publish'
	gormgen.ApplyBasic(model.Publish{}, model.UserProfile{})

	// generate type-safe API for struct 'model.Publish'
	gormgen.ApplyInterface(func(Querier) {}, model.Publish{}, model.UserProfile{})

	gormgen.Execute()
}

func main() {
	generateModelByGen()
}
