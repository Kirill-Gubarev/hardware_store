package service

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
	"address_service/internal/db"
	"strings"
	"github.com/gin-gonic/gin"
	"fmt"
)

type SuggestResponse struct {
	Result struct {
		Items []struct {
			ID       string `json:"id"`
			Name string `json:"full_name"`
			Point struct {
				Lat  float64 `json:"lat"`
				Lon  float64 `json:"lon"`
			} `json:"point"`
		} `json:"items"`
	} `json:"result"`
}

type SuggestItem struct {
	ID        string   `json:"id"`
	Name  string   `json:"name"`
	Lat       float64  `json:"lat"`
	Lon       float64  `json:"lon"`
}

func SuggestAddress(c *gin.Context) {
	query := c.DefaultQuery("q", "")
	if query == "" {
		c.JSON(400, gin.H{"error": "Missing required query parameter 'q'"})
		return
	}
	url := "https://catalog.api.2gis.com/3.0/suggests?key=" + APIKey +
		"&q=" + query +
		"&suggest_type=address&fields=items.point"

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		c.JSON(500, gin.H{"error": "Request timeout or failed"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read response"})
		return
	}

	var suggestResp SuggestResponse
	if err := json.Unmarshal(body, &suggestResp); err != nil {
		c.JSON(500, gin.H{"error": "Failed to parse JSON"})
		return
	}

	var result []SuggestItem
	for _, item := range suggestResp.Result.Items {
		result = append(result, SuggestItem{
			ID:       item.ID,
			Name:     item.Name,
			Lat:      item.Point.Lat,
			Lon:      item.Point.Lon,
		})
	}

	c.JSON(200, result)
}
type DistanceResponse struct {
	Routes []struct {
		Distance int `json:"distance"`
		Duration int `json:"duration"`
	} `json:"routes"`
}
type DistanceItem struct {
	Distance int `json:"distance"`
	Duration int `json:"duration"`
}
func GetDistance(c *gin.Context){
	id1 := c.DefaultQuery("id1", "")
	id2 := c.DefaultQuery("id2", "")
	if id1 == "" || id2 == "" {
		c.JSON(400, gin.H{"error": "Missing required query parameters ids"})
		return
	}
	url := "https://routing.api.2gis.com/get_dist_matrix?key=" + APIKey + "&version=2.0" 

	address1, err := db.GetAddress(id1)
	address2, err2 := db.GetAddress(id2)
	if err != nil || err2 != nil {
		c.JSON(400, gin.H{"error": "Address not found"})
		return
	}
	body := strings.NewReader(fmt.Sprintf(
		`{"points":[
			{"lat": %f,"lon": %f},
			{"lat": %f,"lon": %f}
		],
		"sources": [0],
		"targets": [1]
	}`, *address1.Lat, *address1.Lon, *address2.Lat, *address2.Lon))

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to create request"})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(500, gin.H{"error": "Request timeout or failed"})
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to read response"})
		return
	}

	var distanceResp DistanceResponse
	err = json.Unmarshal(respBody, &distanceResp)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to parse JSON"})
		return
	}
	if len(distanceResp.Routes) == 0 {
		c.JSON(500, gin.H{"error": "No routes found in response"})
		return
	}
	distanceItem := DistanceItem{
		Distance: distanceResp.Routes[0].Distance,
		Duration: distanceResp.Routes[0].Duration,
	}

	c.JSON(200, distanceItem)
}

func GetAddress(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		respondError(c, 400, "Missing address ID")
		return
	}

	address, err := db.GetAddress(id)
	if err != nil {
		respondError(c, 404, "Address not found")
		return
	}
	c.JSON(200, address)
}
func CreateAddress(c *gin.Context) {
	var address db.Address
	if err := c.ShouldBindJSON(&address); err != nil {
		respondError(c, 400, "Invalid JSON")
		return
	}

	id, err := db.CreateAddress(&address)
	if err != nil {
		respondError(c, 400, "Failed to create address")
		return
	}

	address.Id = &id
	c.JSON(201, address)
}
