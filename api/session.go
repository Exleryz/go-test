package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	name := c.Query("name")

	session := sessions.Default(c)
	fmt.Println(session.Get("name"))
	session.Set("name", name)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"session-name": session.Get("name"), "query": name})
}
