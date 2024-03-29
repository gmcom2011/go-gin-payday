package main

import (
	"fmt"
	"os"
	"payday/payday"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		// AllowOrigins:     []string{"payday-e074e.firebaseapp.com"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "payday-e074e.firebaseapp.com"
		// },
		MaxAge: 12 * time.Hour,
	}))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		fmt.Println("response.StatusCode")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/post", func(c *gin.Context) {

		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
	})

	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		dataBody := payday.DataBody(c)
		t := payday.New(dataBody)
		result := t.GetUser(id)
		//fmt.Println("length of result", len(result))

		c.JSON(200, result)
	})

	r.POST("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		dataBody := payday.DataBody(c)
		t := payday.New(dataBody)
		t.UpdateUser(id)

		c.JSON(200, "Update Success.")
	})

	r.POST("/upload/image/:id", func(c *gin.Context) {
		id := c.Param("id")
		fmt.Println("Hello")
		route := payday.App{}
		route.UploadProfile(c.Writer, c.Request, id)

		c.JSON(200, "Upload Success.")
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		payday.DeleteUser(id)

		c.JSON(200, "Delete Success.")
	})

	r.PUT("/user/", func(c *gin.Context) {
		dataBody := payday.DataBody(c)
		t := payday.New(dataBody)
		t.AddUser(dataBody["id"])
		c.JSON(200, "Create User Complete.")
	})

	r.GET("/geturl/:img", func(c *gin.Context) {
		img := c.Param("img")
		result := payday.GetImageUrl(img)
		//fmt.Println("length of result", len(result))

		c.JSON(200, result)
	})
	r.GET("/generatecheckincode/:company", func(c *gin.Context) {
		company := c.Param("company")
		result := payday.GenerateAttendanceCode(company)
		//fmt.Println("length of result", len(result))

		c.JSON(200, result)
	})
	port := os.Getenv("PORT")
	// port := "5000"
	r.Run(":" + port)
}
