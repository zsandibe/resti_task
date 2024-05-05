package v1

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")
	{

		api.GET("/:id")
		// api.POST("/add", h.AddCar)
		// api.DELETE("/delete/:id", h.DeleteCarById)
		// api.GET("/list", h.GetCarsList)
		// api.PUT("/update/:id", h.UpdateCarInfo)

	}
	return router
}
