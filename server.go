package main 

import (
	"net/http"
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/williaminfante/go_test_starter/config"
	"github.com/williaminfante/go_test_starter/migrations"
	"github.com/williaminfante/go_test_starter/controller"
)


var defaultPort = "8080"

func init() {
	godotenv.Load()
}


func main() {
		
	config.ConnectGorm() 
	db := config.GetDb() 
	database, _ := db.DB() 
	defer database.Close() 

	migrations.MigrateTable() 
	port := os.Getenv("DB_PORT")

	if port == "" {
		fmt.Println("the port is blank, falling back to default :8080")
		port = defaultPort
	}

	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

	router.GET("/users", controller.UserInterface.GetAll)
	router.POST("/users", controller.UserInterface.Create)

	host := fmt.Sprintf("%s:", os.Getenv("DB_HOST"))

	err := router.Run(host + defaultPort)
	if err != nil {
		fmt.Printf("Error running server; %s", err)
		return 
	}
}