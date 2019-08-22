package api

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Basic(c *gin.Context) {
	name := c.Query("name")

	session := sessions.Default(c)
	fmt.Println(session.Get("name"))
	session.Set("name", name)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"session-name": session.Get("name"), "query": name})
}

func Multiple(c *gin.Context) {
	name := c.DefaultQuery("name", "defaultName")
	age := c.DefaultQuery("age", "18")

	session := sessions.DefaultMany(c, "session")
	sessionEnv := sessions.DefaultMany(c, os.Getenv("SESSION_NAME"))

	fmt.Printf("name: %v, age: %v\n", session.Get("name"), sessionEnv.Get("age"))

	session.Set("name", name)
	session.Save()
	sessionEnv.Set("age", age)
	sessionEnv.Save()

	c.JSON(http.StatusOK, gin.H{"data": &map[string]string{"name": name, "age": age}})
}
