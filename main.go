package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/config"
	"go-admin/cmd"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:generate swag init --parseDependency --parseDepth=6 --instanceName admin -o ./docs/admin

// @title go-admin API
// @version 2.0.0
// @description 基于Gin + Vue + Element UI的前后端分离权限管理系统的接口文档
// @description 添加qq群: 521386980 进入技术交流群 请先star，谢谢！
// @license.name MIT
// @license.url https://github.com/go-admin-team/go-admin/blob/master/LICENSE.md

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
func main() {

	go cmd.Execute()
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println(sig.String())
		os.Exit(0)
	}()

	time.Sleep(3 * time.Second)
	// 生成静态文件服务
	r := gin.Default()
	// 设置静态文件服务
	dir := "/var/image/"
	if config.ApplicationConfig.Mode == "dev" {
		dir = "D:/images/"
	}
	os.MkdirAll(dir, os.ModeDir)
	r.Static("/images", dir)

	r.Run(":8888")
}
