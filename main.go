package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

// Struktur data User
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

// Handler untuk GET request
func getUser(c echo.Context) error {
	user := User{
		ID:    1,
		Name:  "John Doe",
		Email: "johndoe@example.com",
		Age:   30,
	}

	fmt.Println(user)

	return c.JSON(http.StatusOK, user)
}

// Handler untuk POST request
func createUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON data"})
	}

	response := map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	}

	fmt.Println(response)
	return c.JSON(http.StatusCreated, response)
}

func main() {
	e := echo.New()

	// Define routes
	e.GET("/user", getUser)
	e.POST("/user/create", createUser)

	// Ambil port dari environment variable, default ke 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Jalankan server
	e.Logger.Fatal(e.Start(":" + port))
}
