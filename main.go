package main

import (
	"fmt"
	"log"
	"net/http"

	"explore/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	router := echo.New()

	//CORS
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowCredentials: true,
	}))

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	//Endpoint - Get All Data
	router.GET("/users", controllers.GetAllUsers)

	//Endpoint - Insert Data
	router.POST("/users", controllers.InsertUser)

	//Endpoint - Update Data
	router.PUT("/users/:user_id", controllers.Authenticate(controllers.UpdateUser))

	//Endpoint - Delete Data
	router.DELETE("/users/:user_id", controllers.Authenticate(controllers.DeleteUser))

	//Endpoint - Login User
	router.POST("/user/login", controllers.GetUserLogin)

	//Endpoint - Logout User
	router.POST("/user/logout", controllers.UserLogout)

	http.Handle("/", router)
	fmt.Println("Connected to port 8080")
	log.Println("Connected to port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
