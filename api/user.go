package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-test/serializer"
	"net/http"
)

// GetUserByID 根据id获取用户
func GetUserByID(c *gin.Context) {
	id := c.DefaultQuery("id", "1")
	fmt.Println(id)
	var u serializer.User
	c.Bind(&u)
	fmt.Println(u)
	c.JSON(http.StatusOK, gin.H{"data": u})
}
