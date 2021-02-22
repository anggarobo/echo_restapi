package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// get all users
func getUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "get all users",
		"users":   users,
	})
}

// get user by id
func getUserController(c echo.Context) error {
}

// delete user by id
func deleteUserController(c echo.Context) error {
}

// update user by id
func updateUserController(c echo.Context) error {
}

func createUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}

	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user successfully created",
		"user":    user,
	})
}

func main() {
	e := echo.New()

	e.GET("/users", getUsersController)
	e.GET("/user/:id", getUserController)
	e.POST("/user", createUserController)
	e.PUT("/user/:id", updateUserController)
	e.DELETE("/user/:id", deleteUserController)

	e.Logger.Fatal(e.Start(":8080"))
}
