package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello GO")
	})
	r.POST("/index", func(context *gin.Context) {
		h := gin.H{"test": "ko"}
		context.JSON(http.StatusOK, h)
	})
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/?tn=50000049_hao_pg")
	})

	// 路由组
	vGroup := r.Group("/video")
	{
		vGroup.GET("/move", func(context *gin.Context) {
			h := gin.H{"test": "xx"}
			name := context.Param("name")
			h["name"] = name
			context.JSON(http.StatusOK, h)
		})
		vGroup.GET("/mp3", func(context *gin.Context) {
			h := gin.H{"test": "mp3"}
			context.JSON(http.StatusOK, h)
		})
	}

	r.Run(":9999") // 这边我们使用9999 这个端口
}
