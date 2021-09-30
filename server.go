package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

func main() {
	e := echo.New()
	e.GET("/", GetUsers)
	e.GET("/:id", GetOneUser)
	e.POST("/", PostUser)
	e.PUT("/:id", UpdateUser)
	e.DELETE("/:id", DeleteUser)

	users = []User{
		{Id: 1, Name: "Ivan", Email: "fabrian.ivan@gmail.com", Password: "123456"},
		{Id: 2, Name: "Danu", Email: "danu@gmail.com", Password: "1341rf"},
		{Id: 3, Name: "Pingkan", Email: "pingkan@gmail.com", Password: "131414"},
	}
	e.Logger.Fatal(e.Start(":8080"))
}

func PostUser(c echo.Context) error {
	createUser := User{}
	c.Bind(&createUser)
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success",
		"data":    createUser,
	})
}

func GetUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

func GetOneUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Inputted ID is invalid",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "successful get all users",
		"users":    users[id],
	})
}

func UpdateUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	id -= 1
	if id == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	update := User{}
	c.Bind(&update)
	users[id] = update
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users[id],
	})
}

func DeleteUser(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if id == -1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid id",
		})
	}
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			if i == len(users)-1 {
				users = users[:len(users)-1]
				return c.JSON(http.StatusOK, map[string]interface{}{
					"messages": "successfully delete user",
					"users":    users,
				})
			}
			users = users[i+1:]
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "successfully get all users",
				"users":    users,
			})
		}

	}
	return c.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "invalid id",
	})
}