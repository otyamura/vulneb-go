package main

import (
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
)

type IP struct {
	Address string `json:"address"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("static/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/login/1", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"error": "",
		})
	})

	r.POST("/login/1", func(c *gin.Context) {
		id := c.PostForm("id")
		fmt.Println(id)
		if id == "john" {
			c.HTML(200, "login.html", gin.H{
				"error": "password is incorrect",
			})
		} else {
			c.HTML(200, "login.html", gin.H{
				"error": "user not exist",
			})
		}
	})

	r.GET("/login/2", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"error": "",
		})
	})

	r.POST("/login/2", func(c *gin.Context) {
		id := c.PostForm("id")
		fmt.Println(id)
		if id == "john" {
			time.Sleep(3 * time.Second)
		}
		c.HTML(200, "login.html", gin.H{
			"error": "user or password is incorrect",
		})
	})

	r.POST("/ping", func(c *gin.Context) {
		var param IP
		if err := c.Bind(&param); err != nil {
			c.JSON(400, gin.H{"message": "Invalid parameter"})
			return
		}

		commnd := "ping -c 1 -W 1 " + param.Address + " 1>&2"
		result, _ := exec.Command("sh", "-c", commnd).CombinedOutput()

		c.JSON(200, gin.H{
			"result": string(result),
		})
	})

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
