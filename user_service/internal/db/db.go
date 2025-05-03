package db

import (
	"github.com/jackc/pgx/v5"
	"user_service/internal/utils"
	"context"
)

var conn *pgx.Conn

func Connect() (error) {
	cfg, err := utils.GetDatabaseConfig()
	if err != nil{
		return err
	}
	connection_str := "postgresql://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName
	conn, err = pgx.Connect(context.Background(), connection_str)
	return err
}

func Close() (error) {
	if conn != nil {
		return conn.Close(context.Background())
	}
	return nil
}
