package db

import (
	"context"
	"fmt"
)

func GetUser(id string) (*User, error){
	query := `
	SELECT
		users.id,
		users.login,
		roles.id,
		roles.name
	FROM users
	LEFT JOIN roles ON users.role_id = roles.id
	WHERE users.id = $1;`

	user := NewUser()
	row := conn.QueryRow(context.Background(), query, id)
	err := row.Scan(
		&user.Id,
		&user.Login,
		&user.Role.Id,
		&user.Role.Name)
 
	return &user, err
}

func GetUsers(limit, offset int) (users []User, err error) {
	query := `
	SELECT
		users.id,
		users.login,
		roles.id,
		roles.name
	FROM users
	LEFT JOIN roles ON users.role_id = roles.id
    LIMIT $1 OFFSET $2;`

	rows, err := conn.Query(context.Background(), query, limit, offset)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		user := NewUser()
		err = rows.Scan(
			&user.Id, 
			&user.Login, 
			&user.Role.Id,
			&user.Role.Name)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, rows.Err()
}

func validRoleId(role *Role) (any) { 
	if role == nil {
		return nil
	}
	return role.Id
}
func CreateUser(user *User) (id string, err error) {
	query := `
	INSERT INTO users (login, password, role_id)
	VALUES ($1, crypt($2, password), $3)
	RETURNING id;`

	err = conn.QueryRow(context.Background(), query,
		user.Login, 
		user.Password, 
		validRoleId(user.Role)).Scan(&id)
			
	return id, err
}

func DeleteUser(id string) error {
	query := `
	DELETE FROM users
	WHERE id = $1;`

	result, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}
	return nil
}
