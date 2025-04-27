package db

type Unit struct{
	Id    *string  `json:"id"`
	Name  *string  `json:"name"`
}

type Product struct{
	Id              *string  `json:"id"`
	Name            *string  `json:"name"`
	Description     *string  `json:"description"`
	Unit            *Unit    `json:"unit"`
	ImageId         *string  `json:"imageId"`
	ManufacturerId  *string  `json:"manufacturerId"`
}
