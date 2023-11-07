package main

import (
	"github.com/1Nelsonel/corbiz/database"
	"github.com/1Nelsonel/corbiz/router"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDB()
}

func main()  {

	// Database connect
	sqlDb, err := database.DBConn.DB()

	if err != nil {
		panic("Error in connection")
	}

	defer sqlDb.Close()
	

	app := gin.Default()

	// Templates
	app.LoadHTMLGlob("views/*")

		// Serve static files (CSS, JavaScript, images) from the "public" directory
	app.Static("/static", "./public")


	// Initialize routes by calling SetupRoutes
	router.SetupRoutes(app)

	app.Run()
}