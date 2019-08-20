package server

import (
	"go-test/api"

	"github.com/gin-gonic/gin"
	//"github.com/thinkerou/favicon"
	"net/http"
)

// Router 路由配置
func Router() *gin.Engine {
	r := gin.Default()
	//r.Use(favicon.New("./favicon.ico")) // 运行目录下的favicon.ico
	// 路由
	group := r.Group("/")
	{
		// 测试服务器是否存活
		group.GET("ping", api.Ping)
		group.GET("set/cookie", api.SetCookie)
		group.GET("get/header", api.GetHeader)
		group.GET("get/body", api.GetBody)
		group.GET("get/jsonp", api.Jsonp)

		// Listen and serve on 0.0.0.0:8080
		r.Run(":8080")

		// 此规则能够匹配/user/john这种格式，但不能匹配/user/ 或 /user这种格式
		group.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello %s", name)
		})
		// 但是，这个规则既能匹配/user/john/格式也能匹配/user/john/send这种格式
		// 如果没有其他路由器匹配/user/john，它将重定向到/user/john/
		group.GET("/user/:name/*action", func(c *gin.Context) {
			name := c.Param("name")
			action := c.Param("action")
			message := name + " is " + action
			c.String(http.StatusOK, message)
		})
		// 匹配的url格式:  /welcome?firstName=Jane&lastName=Doe
		group.GET("/welcome", func(c *gin.Context) {
			firstName := c.DefaultQuery("firstName", "Guest")
			lastName := c.Query("lastName") // 是 c.Request.URL.Query().Get("lastname") 的简写

			c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
		})
		group.POST("/form_post", func(c *gin.Context) {
			message := c.PostForm("message")
			nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

			c.JSON(200, gin.H{
				"status":  "posted",
				"message": message,
				"nick":    nick,
			})
		})

		group.POST("get/user", api.GetUserByID)

	}
	gin.SetMode("debug")
	return r
}
