package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context)  {
	c.HTML(http.StatusOK, "home.html", c)
}

func About(c *gin.Context)  {
	c.HTML(http.StatusOK, "about.html", c)
}