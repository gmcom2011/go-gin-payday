package main

import (
	"encoding/json"
	"fmt"
	"os"
	"payday/payday"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
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
		dataBody := dataBody(c)
		t := payday.New(dataBody)
		result := t.GetUser(id)
		fmt.Println("length of result", len(result))

		c.JSON(200, result)
	})

	r.POST("/user/:v", func(c *gin.Context) {
		id := c.Param("id")
		dataBody := dataBody(c)
		t := payday.New(dataBody)
		result := t.UpdateUser(id)

		c.JSON(200, result)
	})

	r.DELETE("/user/:v", func(c *gin.Context) {
		buf := make([]byte, 1024)
		body, _ := c.Request.Body.Read(buf)
		//reqBody := string(buf[0:body])
		reqBody2 := buf[0:body]
		//fmt.Println(reflect.TypeOf(reqBody))
		var reqMap map[string]interface{}
		json.Unmarshal(reqBody2, &reqMap)
		fmt.Println(string(reqBody2))
		fmt.Println(reqMap)

		c.JSON(200, string(reqBody2))
	})

	r.PUT("/user/", func(c *gin.Context) {
		dataBody := dataBody(c)
		t := payday.New(dataBody)
		fmt.Println(t.FirstNameTh)
		t.AddUser(dataBody["id"])
		fmt.Println("response.StatusCode")
		c.JSON(200, "Create User Complete.")
	})
	port := os.Getenv("PORT")
	r.Run(":" + port)
}

func dataBody(c *gin.Context) map[string]string {
	buf := make([]byte, 1024)
	rawBody, _ := c.Request.Body.Read(buf)
	jsonBody := buf[0:rawBody]
	fmt.Println(jsonBody)
	var mapBody map[string]string
	json.Unmarshal(jsonBody, &mapBody)
	return mapBody
}
