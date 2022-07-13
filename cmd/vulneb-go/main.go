package main

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Account struct {
	ID       string
	Password string
}

func main() {
	r := gin.Default()

	db, err := sql.Open("mysql", "root:root@tcp(db:3306)/vulneb")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

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
		if id == "Amelia" {
			c.HTML(200, "login.html", gin.H{
				"error": "password is incorrect",
			})
		} else {
			c.HTML(200, "login.html", gin.H{
				"error": "id not exist",
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
		if id == "Liam" {
			time.Sleep(2 * time.Second)
		}
		c.HTML(200, "login.html", gin.H{
			"error": "id or password is incorrect",
		})
	})

	r.GET("/login/3", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"error": "",
		})
	})

	r.GET("/login/4", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"error": "",
		})
	})

	r.POST("/login/4", func(c *gin.Context) {
		id := c.PostForm("id")
		pass := c.PostForm("password")
		fmt.Println(id, pass)
		query := "SELECT * FROM accounts WHERE id=? AND password=?"
		stmt, err := db.Prepare(query)
		if err != nil {
			c.HTML(200, "login.html", gin.H{
				"error": err,
				"query": query,
			})
			return
		}
		rows, err := stmt.Query(id, pass)
		if err != nil {
			c.HTML(200, "login.html", gin.H{
				"error": err,
				"query": query,
			})
			return
		}
		defer rows.Close()

		var accounts []Account
		if !rows.Next() {
			c.HTML(200, "login.html", gin.H{
				"error": "no rows in result set",
				"query": query,
			})
			return
		} else {
			var a Account
			err := rows.Scan(&a.ID, &a.Password)
			if err != nil {
				log.Fatal(err)
			}
			accounts = append(accounts, a)
		}
		for rows.Next() {
			var a Account
			err := rows.Scan(&a.ID, &a.Password)
			if err != nil {
				log.Fatal(err)
			}
			accounts = append(accounts, a)
		}
		c.HTML(200, "success.html", gin.H{
			"accounts": accounts,
		})
	})

	r.POST("/login/3", func(c *gin.Context) {
		id := c.PostForm("id")
		pass := c.PostForm("password")
		fmt.Println(id, pass)
		query := "SELECT * FROM accounts WHERE id='" + id + "' AND password='" + pass + "'"
		rows, err := db.Query(query)
		if err != nil {
			c.HTML(200, "login.html", gin.H{
				"error": err,
				"query": query,
			})
			return
		}
		defer rows.Close()

		var accounts []Account
		if !rows.Next() {
			c.HTML(200, "login.html", gin.H{
				"error": "no rows in result set",
				"query": query,
			})
			return
		} else {
			var a Account
			err := rows.Scan(&a.ID, &a.Password)
			if err != nil {
				log.Fatal(err)
			}
			accounts = append(accounts, a)
		}
		for rows.Next() {
			var a Account
			err := rows.Scan(&a.ID, &a.Password)
			if err != nil {
				log.Fatal(err)
			}
			accounts = append(accounts, a)
		}
		c.HTML(200, "success.html", gin.H{
			"accounts": accounts,
		})
	})

	r.GET("/tldr", func(c *gin.Context) {
		c.HTML(200, "tldr.html", nil)
	})
	r.POST("/tldr", func(c *gin.Context) {
		name := c.PostForm("name")
		cmd := "tldr " + name
		remove_color := `'| sed -e 's/\x1b\[[0-9;]*m//g' 1>&2`
		result, _ := exec.Command("sh", "-c", cmd).CombinedOutput()
		removed, _ := exec.Command("sh", "-c", "echo '"+string(result)+remove_color).CombinedOutput()

		c.HTML(200, "tldr.html", gin.H{
			"command": cmd,
			"result":  string(removed),
		})
	})
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
