package main

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"os"
)

func GetHost(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
	}
	c.JSON(200, gin.H{
		"hostname": hostname,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", GetHost)
	r.POST("/", func(c *gin.Context) {
		body, _ := ioutil.ReadAll(c.Request.Body)
		log.Println(string(body))
	})
	r.Run(":3000")
}
