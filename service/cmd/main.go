package main

import (
	"fmt"
	"h5/config"
	"h5/handlers"
	"h5/repository"
	"h5/router"
	"h5/service"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	// 初始化数据库
	db, err := gorm.Open(mysql.Open(config.GlobalConfig.Database.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 初始化repository
	warrantyRepo := repository.NewWarrantyCardRepo(db)
	institutionRepo := repository.NewInstitutionRepo(db) // 确保这行存在
	// 初始化service
	warrantyService := service.NewWarrantyCardService(warrantyRepo)
	institutionService := service.NewInstitutionService(institutionRepo) // 确保这行存在

	// 初始化handlers
	handlers.InitHandlers(warrantyService)
	handlers.InitInstitutionHandlers(institutionService) // 确保这行存在

	// 设置路由
	r := router.SetupRouter()

	// 启动服务器
	serverConfig := config.GlobalConfig.Server
	log.Printf("Starting server on port %d in %s mode", serverConfig.Port, serverConfig.Mode)
	if err := r.Run(fmt.Sprintf(":%d", serverConfig.Port)); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}