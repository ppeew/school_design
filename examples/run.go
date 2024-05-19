//go:build examples
// +build examples

package main

import (
	"fmt"
	"github.com/go-admin-team/go-admin-core/sdk"
	"go-admin/app/admin/models"
	cModels "go-admin/common/models"
	"gorm.io/gorm"
	"time"

	myCasbin "github.com/go-admin-team/go-admin-core/sdk/pkg/casbin"
	"gorm.io/driver/mysql"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:518888@tcp(139.159.234.134:3306)/go_admin?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	syncEnforce := myCasbin.Setup(db, "sys_")
	sdk.Runtime.SetDb("*", db)
	sdk.Runtime.SetCasbin("*", syncEnforce)

	// 生成Model
	db.AutoMigrate(&models.Blogs{}, &models.Collects{}, &models.Comments2{})

	data := models.Blogs{
		Model:    cModels.Model{},
		Username: "555",
		Msg:      "你好啊，测试",
		Star:     "1",
		Collect:  "2",
		ModelTime: cModels.ModelTime{
			CreatedAt: time.Now(),
		},
	}
	err = db.Create(&data).Error
	if err != nil {
		fmt.Sprintf("BlogsService Insert error:%s \r\n", err)
	}

	//e := gin.Default()
	//sdk.Runtime.SetEngine(e)
	//log.Fatal(e.Run(":9999"))
}
