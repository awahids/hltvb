package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouters(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	validate := validator.New()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/helloworld")
		{
			eg.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "Hello, World!"})
			})
		}

		BookRouter(v1, validate, db)
	}

	return r
}
