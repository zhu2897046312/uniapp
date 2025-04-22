package router

import (
	"h5/handlers"
	"h5/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.Logger(), middleware.Recovery())
	r.Use(middleware.Cors())
	r.Static("/uploads", "./uploads")
	api := r.Group("/api")
	{
		api.POST("/login",handlers.Login)
		// 保修卡相关路由
		warranty := api.Group("/warranty")
		{
			warranty.POST("/create", handlers.CreateWarrantyCard)
			warranty.GET("/list", handlers.GetWarrantyCardList)
			warranty.GET("/:id",handlers.GetWarrantyCardByID)
			warranty.POST("/upload", handlers.UploadImage)
			warranty.GET("/info",handlers.GetWarrantyCardByNumber)
		}

		institution := api.Group("/institution")
		{
			institution.GET("/info", handlers.GetInstitutionByToken)
			institution.POST("/update-password", handlers.UpdatePassword)
		}
		
	}

	return r
}