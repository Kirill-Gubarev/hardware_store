package db

type Address struct{
	Id     *string   `json:"id"`
	IdAPI  *string   `json:"idAPI"`
	Name   *string   `json:"name"`
	Lat    *float64  `json:"lat"`
	Lon    *float64  `json:"lot"`
}
