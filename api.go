package main

import "github.com/gin-gonic/gin"

func main() {
	Storage := NewMemoryStorage()
	handler := NewHandler(Storage)
	router := gin.Default()

	router.GET("/employee/:id", handler.GetEmployee)
	router.DELETE("/employee/:id", handler.DeleteEmployee)
	router.POST("/employee", handler.CreateEmployee)
	router.PUT("/employee/:id", handler.UpdateEmployee)

	router.Run(":8080")

}
