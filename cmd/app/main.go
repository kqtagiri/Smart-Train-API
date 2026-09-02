package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "HELLO FROM GIN")
	})

	r.Run(":9111")
	fmt.Println("Программа завершена")

}
