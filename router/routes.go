package router

import (
	"github.com/1Nelsonel/corbiz/pkg"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine)  {
	r.GET("/", pkg.Home)
    r.GET("about/", pkg.About)
}