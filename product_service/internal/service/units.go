package service 

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"product_service/internal/db"
)

func GetUnit(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		respondError(c, 400, "Missing unit ID")
		return
	}

	unit, err := db.GetUnit(id)
	if err != nil {
		respondError(c, 404, "Unit not found")
		return
	}
	c.JSON(200, unit)
}

func GetUnits(c *gin.Context) {
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

	units, err := db.GetUnits(limit, offset)
	if err != nil {
		respondError(c, 404, "Units not found")
		return
	}

	if len(units) == 0 {
		c.JSON(200, []any{})
		return
	}
	c.JSON(200, units)
}

func CreateUnit(c *gin.Context) {
	var unit db.Unit
	if err := c.ShouldBindJSON(&unit); err != nil {
		respondError(c, 400, "Invalid JSON")
		return
	}

	id, err := db.CreateUnit(&unit)
	if err != nil {
		respondError(c, 400, "Failed to create unit")
		return
	}

	unit.Id = &id
	c.JSON(201, unit)
}

func DeleteUnit(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		respondError(c, 400, "Missing unit ID")
		return
	}

	err := db.DeleteUnit(id)
	if err != nil {
		respondError(c, 404, "Unit not found or already deleted")
		return
	}

	respondSuccess(c, 200, "Unit deleted")
}
