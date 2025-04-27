package db

import (
	"context"
	"fmt"
)

func GetProduct(id string) (*Product, error){
	query := `
	SELECT
		products.id,
		products.name,
		products.description,
		products.image_id,
		products.manufacturer_id,
		units.id,
		units.name
	FROM products
	LEFT JOIN units ON products.unit_id = units.id
	WHERE products.id = $1;`

	product := &Product{Unit: &Unit{}}
	row := conn.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&product.Id, 
		&product.Name, 
		&product.Description, 
		&product.ImageId, 
		&product.ManufacturerId,
		&product.Unit.Id,
		&product.Unit.Name)
 
	return product, err
}

func GetProducts(limit, offset int) (products []Product, err error) {
	query := `
	SELECT
		products.id,
		products.name,
		products.description,
		products.image_id,
		products.manufacturer_id,
		units.id,
		units.name
	FROM products
	LEFT JOIN units ON products.unit_id = units.id
    LIMIT $1 OFFSET $2;`

	rows, err := conn.Query(context.Background(), query, limit, offset)
	if err != nil {
		return products, err
	}
	defer rows.Close()

	for rows.Next() {
		product := Product{Unit: &Unit{}}
		err = rows.Scan(
			&product.Id, 
			&product.Name, 
			&product.Description, 
			&product.ImageId, 
			&product.ManufacturerId,
			&product.Unit.Id,
			&product.Unit.Name)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, rows.Err()
}

func validUnitId(unit *Unit) (any) { 
	if unit == nil {
		return nil
	}
	return unit.Id
}
func CreateProduct(product *Product) (id string, err error) {
	query := `
	INSERT INTO products
	VALUES (DEFAULT, $1, $2, $3, $4, $5)
	RETURNING id;`

	err = conn.QueryRow(context.Background(), query,
		product.Name, 
		product.Description, 
		validUnitId(product.Unit),
		product.ImageId, 
		product.ManufacturerId).Scan(&id)
			
	return id, err
}

func DeleteProduct(id string) error {
	query := `
	DELETE FROM products
	WHERE id = $1;`

	result, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("product with id %s not found", id)
	}
	return nil
}
