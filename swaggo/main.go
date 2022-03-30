package main

import (
	"net/http"

	_ "github.com/diwangtseb/essays/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://www.topgoer.com

// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1

func main() {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", HandleHello)
		// v1.POST("/login", HandleLogin)
		// v1Auth := r.Use(HandleAuth)
		// {
		//     v1Auth.POST("/upload", HandleUpload)
		//     v1Auth.GET("/list", HandleList)
		// }
	}

	r.Run(":8080")
}

// @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept json
// @Param who query string true "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /hello [get]
func HandleHello(c *gin.Context) {
	who := c.Query("who")

	if who == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "who are u?"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "hello " + who})
}
