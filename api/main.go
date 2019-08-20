package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/serializer"
	"io/ioutil"
	"net/http"
)

// Ping ping
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, serializer.Response{
		Status: 0,
		Msg:    "pong",
	})

}

// SetCookie 设置cookie
func SetCookie(c *gin.Context) {
	// cookie test
	// name value maxAge(秒, 过期时间, 0 -> session) path domain secure httpOnly
	c.SetCookie("name", "value", 0, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"msg": "set cookie"})
}

// GetHeader 获取请求头
func GetHeader(c *gin.Context) {
	fmt.Println("---header: start")
	for key, value := range c.Request.Header {
		fmt.Printf("%v : %v\n", key, value)
	}
	fmt.Println("---header: end")
	c.JSON(http.StatusOK, gin.H{"msg": "getHeader success", "data": c.Request.Header})
}

// GetBody 获取请求body
func GetBody(c *gin.Context) {
	fmt.Println("---body: start")
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "body获取失败"})
		return
	}
	fmt.Printf("%s\n", body)
	fmt.Println("---body: end")
	c.JSON(http.StatusOK, gin.H{"msg": "getBody success", "data": body})
}

// GetFormData
// JSONP
func Jsonp(c *gin.Context) {
	data := map[string]interface{}{
		"msg": "test",
	}

	// 访问 http://localhost:8080/get/jsonp?callback=call
	// 将会输出:   call({"msg":"test"})
	c.JSONP(http.StatusOK, data)
}
