package sql

import (
	"database/sql"
	"errors"
	"fmt"
)

func SetValuesForDB(port string, user string, password string, dbname string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=disable", user, password, dbname, port)
}

func GetDB() (*sql.DB, error) {
	var err error
	var db *sql.DB

	//connStr := SetValuesForDB("postgres", "postgres", "postgres", "5432")
	db, err = sql.Open("postgres", "user=postgres password=qwerty dbname=postgres port=5436 sslmode=disable")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func GetAccessLogin(db *sql.DB, login string, password string) (bool, error) {
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

func AddOrderToDB(db *sql.DB, order_name string, name string, surname string, email string, adress string) (bool, error) {
	rows, err := db.Query("select is_not_full();")

	if err != nil {
		return false, err
	}
	var try bool
	_ = rows.Scan(&try)

	if try == false {
		return try, errors.New("Нет места на складе")
	}

	str := fmt.Sprintf("select  insert_product('%s', '%s', '%s', '%s', '%s');", order_name, name, surname, adress, email)
	obj, err := db.Exec(str)
	if err != nil {
		return false, err
	}
	fmt.Println(obj.LastInsertId())

	return true, nil
}
