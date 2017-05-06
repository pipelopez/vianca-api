package main

import (
  "net/http"
  "os"
  "gopkg.in/gin-gonic/gin.v1"
)

func main() {
  port := os.Getenv("PORT")
  router := gin.Default()

  if port == "" {
    port = "8080"
  }

  router.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{ "message": "hola mundo" })
  })

  router.Run(":" + port)
}