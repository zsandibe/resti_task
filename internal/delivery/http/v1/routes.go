package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zsandibe/resti_task/docs"
)

func (h *Handler) Routes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := router.Group("/api/v1")
	{

		account := api.Group("/accounts")
		{
			account.POST("/create", h.CreateAccount)
			account.GET("/", h.GetAllAccountsWithTransactions)
		}

		transaction := api.Group("/transactions")
		{
			transaction.POST("/create", h.CreateTransaction)
			transaction.GET("/:id", h.GetAllTransactionsByAccountId)
		}

	}
	return router
}
