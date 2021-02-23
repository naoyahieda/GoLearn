package second

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func second() {
    engine:= gin.Default()
    ua := ""
     // ミドルウェアを使用
    engine.Use(func(c *gin.Context) {
        ua = c.GetHeader("User-Agent")
        c.Next()
    })
    engine.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message":    "hello world",
            "User-Agent": ua,
        })
    })
    engine.Run(":3000")
}