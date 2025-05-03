package db

import (
	"context"
	"fmt"
)

func GetRole(id string) (*Role, error){
	query := `SELECT * FROM roles WHERE id = $1;`

	role := NewRole()
	row := conn.QueryRow(context.Background(), query, id)
	err := row.Scan(&role.Id, &role.Name)

	return &role, err
}

func GetRoles(limit, offset int) (roles []Role, err error) {
	query := `
    SELECT *
    FROM roles
    LIMIT $1 OFFSET $2;`

	rows, err := conn.Query(context.Background(), query, limit, offset)
	if err != nil {
		return roles, err
	}
	defer rows.Close()

	for rows.Next() {
		role := NewRole()
		err = rows.Scan(&role.Id, &role.Name)
		if err != nil {
			return roles, err
		}
		roles = append(roles, role)
	}

	return roles, rows.Err()
}

func CreateRole(role *Role) (id string, err error) {
	query := `
	INSERT INTO roles
	VALUES (DEFAULT, $1)
	RETURNING id;`

	err = conn.QueryRow(context.Background(), query, role.Name).Scan(&id)
	return id, err
}

func DeleteRole(id string) error {
	query := `
	DELETE FROM roles
	WHERE id = $1;`

	result, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("role with id %s not found", id)
	}
	return nil
}
