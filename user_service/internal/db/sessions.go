package db

import (
	"crypto/rand"
	"encoding/base64"
	"context"
	"fmt"
	"log"
)

func GetSession(login, token string) (*Session, error){
	query := `
	SELECT
		sessions.id,
		sessions.token,
		sessions.created_at,
		users.id,
		users.login,
		roles.id,
		roles.name
	FROM sessions
	LEFT JOIN users ON sessions.user_id = users.id
	LEFT JOIN roles ON users.role_id = roles.id
	WHERE sessions.token = $2 AND users.login = $1;`

	session := NewSession()
	row := conn.QueryRow(context.Background(), query, login, token)
	err := row.Scan(
		&session.Id,
		&session.Token,
		&session.CreatedAt,
		&session.User.Id,
		&session.User.Login,
		&session.User.Role.Id,
		&session.User.Role.Name)
 
	return &session, err
}

func generateToken() (string, error) {
	b := make([]byte, 128)
	_, err := rand.Read(b);
	if  err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

func CreateSession(login, password string) (token string, err error) {
	query := `
	INSERT INTO sessions(user_id, token)
	SELECT id, $3 FROM users
	WHERE login=$1 AND password=crypt($2,password)`

	token, err = generateToken()
	if err != nil {
		log.Println(err)
		return "", err
	}

	_, err = conn.Exec(context.Background(), query, login, password, token)
	return token, err
}
func UpdateSession(id string) (err error) {
	query := `UPDATE sessions SET created_at = NOW() WHERE sessions.id = $1`

	_, err = conn.Exec(context.Background(), query, id)
 
	return err
}

func DeleteSession(id string) error {
	query := `
	DELETE FROM sessions
	WHERE id = $1;`

	result, err := conn.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("session with id %s not found", id)
	}
	return nil
}
