package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/matheusjost/go-bettings-api/configs"
)

func OpenConn() (*sql.DB, error) {
	conf := configs.GetDBConfig()
	sconn := fmt.Sprintf(
		"port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Port,
		conf.Host,
		conf.User,
		conf.Pass,
		conf.Database)

	conn, err := sql.Open("postgres", sconn)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	return conn, err
}
