package db

import (
	"context"
	"fmt"
)

func GetAddress(id string) (*Address, error){
	query := `SELECT * FROM addresses WHERE id = $1;`

	address := &Address{}
	row := conn.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&address.Id,
		&address.IdAPI,
		&address.Name,
		&address.Lat,
		&address.Lon)

	return address, err
}
func CreateAddress(address *Address) (id string, err error) {
	query := `
	INSERT INTO addresses
	VALUES (DEFAULT, $1, $2, $3, $4)
	RETURNING id;`

	err = conn.QueryRow(context.Background(), query, 
		address.IdAPI,
		address.Name,
		address.Lat,
		address.Lon).Scan(&id)
	return id, err
}

func DeleteAddress(id string) error {
	query := `
	DELETE FROM addresses
	WHERE id = $1;`

	result, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("address with id %s not found", id)
	}
	return nil
}

