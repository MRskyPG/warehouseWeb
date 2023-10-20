package sql

import (
	"database/sql"
	"fmt"
)

func SetValuesForDB(port string, user string, password string, dbname string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbname, port)
}

func GetDB() (*sql.DB, error) {
	var err error
	var db *sql.DB

	//connStr := SetValuesForDB("postgres", "postgres", "postgres", "5432")
	db, err = sql.Open("postgres", "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAccess(db *sql.DB, login string, password string) (bool, error) {
	var condition bool
	str := fmt.Sprintf("select * from CHECKLOGIN('%s', '%s')", login, password)
	rows, err := db.Query(str)
	if err != nil {
		return false, err
	}

	for rows.Next() {
		_ = rows.Scan(&condition)
	}

	return condition, nil
}
