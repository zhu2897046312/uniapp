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

	api := r.Group("/api")
	{
		// 保修卡相关路由
		warranty := api.Group("/warranty")
		{
			warranty.POST("/create", handlers.CreateWarrantyCard)
			warranty.GET("/list", handlers.GetWarrantyCardList)
			warranty.GET("/:id",handlers.GetWarrantyCardByID)
			warranty.POST("/upload", handlers.UploadImage)
		}

		institution := api.Group("/institution")
		{
			institution.GET("/info", handlers.GetInstitutionByToken)
		}
		
	}

	return r
}