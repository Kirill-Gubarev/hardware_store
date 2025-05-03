package service 

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"user_service/internal/db"
)

func GetRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		respondError(c, 400, "Missing role ID")
		return
	}

	role, err := db.GetRole(id)
	if err != nil {
		respondError(c, 404, "Role not found")
		return
	}
	c.JSON(200, role)
}

func GetRoles(c *gin.Context) {
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

	roles, err := db.GetRoles(limit, offset)
	if err != nil {
		respondError(c, 404, "Roles not found")
		return
	}

	if len(roles) == 0 {
		c.JSON(200, []any{})
		return
	}
	c.JSON(200, roles)
}

func CreateRole(c *gin.Context) {
	var role db.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		respondError(c, 400, "Invalid JSON")
		return
	}

	id, err := db.CreateRole(&role)
	if err != nil {
		respondError(c, 400, "Failed to create role")
		return
	}

	role.Id = &id
	c.JSON(201, role)
}

func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		respondError(c, 400, "Missing role ID")
		return
	}

	err := db.DeleteRole(id)
	if err != nil {
		respondError(c, 404, "Role not found or already deleted")
		return
	}

	respondSuccess(c, 200, "Role deleted")
}
