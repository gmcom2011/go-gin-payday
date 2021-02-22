package main

import (
	"encoding/json"
	"fmt"
	"net/http"
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
		t := payday.New("", "", "", "", "", "", "", "")
		result := t.GetUser(id)

		c.JSON(200, result)
	})

	r.POST("/user/:v", func(c *gin.Context) {
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
		buf := make([]byte, 1024)

		rawBody, _ := c.Request.Body.Read(buf)
		jsonBody := buf[0:rawBody]
		fmt.Println(jsonBody)
		var mapBody map[string]string
		json.Unmarshal(jsonBody, &mapBody)
		fmt.Println(mapBody)

		t := payday.New(mapBody["firstNameEn"], mapBody["lastNameEn"], mapBody["firstNameTh"], mapBody["lastNameth"], mapBody["titleEn"], mapBody["titleTh"], mapBody["displayName"], mapBody["userType"])
		fmt.Println(t.FirstNameTh)
		t.AddUser(mapBody["id"])
		fmt.Println("response.StatusCode")
		//fmt.Println(result)
		c.JSON(200, "Create User Complete.")
	})

	r.GET("/someDataFromReader", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		fmt.Println(response.StatusCode)
		if err != nil || response.StatusCode != http.StatusOK {
			fmt.Println(err)
			c.Status(http.StatusServiceUnavailable)
			return
		}

		reader := response.Body
		defer reader.Close()
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
