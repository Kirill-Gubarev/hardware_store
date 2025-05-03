package service 

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"user_service/internal/db"
	"fmt"
)

type UserResponse struct{
	Id        *string  `json:"id"`
	Login     *string  `json:"login"`
	Role      *db.Role `json:"role"`
}
func ConvertUserToUserResponse(user *db.User)UserResponse{
	return UserResponse{
		Id:     user.Id,
		Login:  user.Login,
		Role:   user.Role,
	}
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		respondError(c, 400, "Missing user ID")
		return
	}

	var user, err = db.GetUser(id)
	if err != nil {
		respondError(c, 404, "User not found")
		return
	}
	c.JSON(200, ConvertUserToUserResponse(user))
}

func GetUsers(c *gin.Context) {
	offset, err1 := strconv.Atoi(c.DefaultQuery("offset", "0"))
	limit, err2 := strconv.Atoi(c.DefaultQuery("limit", "20"))

	if err1 != nil || err2 != nil {
		respondError(c, 400, "Offset and limit must be integers")
		return
	}
	if limit < 0 || offset < 0 {
		respondError(c, 400, "Offset and limit must be non-negative")
		return
	}

	limit = min(limit, 200)
	offset *= limit

	users, err := db.GetUsers(limit, offset)
	if err != nil {
		fmt.Println(err)
		respondError(c, 404, "Users not found")
		return
	}

	if len(users) == 0 {
		c.JSON(200, []any{})
		return
	}

	response := make([]UserResponse, len(users))
	for i, v := range users{
		response[i] = ConvertUserToUserResponse(&v)
	}
	c.JSON(200, response)
}

func CreateUser(c *gin.Context) {
	var user db.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		respondError(c, 400, "Invalid JSON")
		return
	}

	id, err := db.CreateUser(&user)
	if err != nil {
		respondError(c, 400, "Failed to create user")
		return
	}

	user.Id = &id
	c.JSON(201, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		respondError(c, 400, "Missing user ID")
		return
	}

	err := db.DeleteUser(id)
	if err != nil {
		respondError(c, 404, "User not found or already deleted")
		return
	}

	respondSuccess(c, 200, "User deleted")
}
