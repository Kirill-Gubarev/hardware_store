package db

import (
	"context"
	"fmt"
)

func GetUnit(id string) (*Unit, error){
	query := `SELECT * FROM units WHERE id = $1;`

	unit := &Unit{}
	row := conn.QueryRow(context.Background(), query, id)
	err := row.Scan(&unit.Id, &unit.Name)

	return unit, err
}

func GetUnits(limit, offset int) (units []Unit, err error) {
	query := `
    SELECT *
    FROM units
    LIMIT $1 OFFSET $2;`

	rows, err := conn.Query(context.Background(), query, limit, offset)
	if err != nil {
		return units, err
	}
	defer rows.Close()

	for rows.Next() {
		unit := Unit{}
		err = rows.Scan(&unit.Id, &unit.Name)
		if err != nil {
			return units, err
		}
		units = append(units, unit)
	}

	return units, rows.Err()
}

func CreateUnit(unit *Unit) (id string, err error) {
	query := `
	INSERT INTO units
	VALUES (DEFAULT, $1)
	RETURNING id;`

	err = conn.QueryRow(context.Background(), query, unit.Name).Scan(&id)
	return id, err
}

func DeleteUnit(id string) error {
	query := `
	DELETE FROM units
	WHERE id = $1;`

	result, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("unit with id %s not found", id)
	}
	return nil
}
