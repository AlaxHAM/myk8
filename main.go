package main

import (
  "mygin/initiallize"
  "mygin/router"
)

func main() {
  // engine := gin.Default()

  // engine.GET("/", func(c *gin.Context) {
  //     c.JSON(200, gin.H{
  //         "message": "Hello, World!",
  //     })
  // })
  // engine.Run("127.0.0.1:8080")

  // RouteGroup(engine)
  // _ = engine.Run("127.0.0.1:8080")

  engine := router.InitRouter()
  initiallize.InitK8SConnect()
  _ = engine.Run("127.0.0.1:8080")
}
