package sql

import (
	"database/sql"
	"errors"
	"fmt"
	Utils "warehouseWeb/internal/searchStruct"
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

	if try == true {
		return try, errors.New("Нет места на складе")
	}

	str := fmt.Sprintf("call insert_product('%s', '%s', '%s', '%s', '%s');", order_name, name, surname, adress, email)
	obj, err := db.Exec(str)
	if err != nil {
		return false, err
	}
	fmt.Println(obj.LastInsertId())

	return true, nil
}

func correctInputArg(arg string, sql_argname string, need_delim bool) (string, bool) {
	if arg != "" {
		arg = sql_argname + ":=" + "'" + arg + "'"
		if need_delim {
			arg = ", " + arg
		}
		need_delim = true
	}
	return arg, need_delim
} 
func createSelectStr(order_name string, cl_name string, cl_surname string, email string, dp_address string) (string) {
	var need_delim bool = false
	order_name, need_delim = correctInputArg(order_name, "prod_name", need_delim)
	cl_name, need_delim = correctInputArg(cl_name, "cl_name", need_delim)
	cl_surname, need_delim = correctInputArg(cl_surname, "cl_surname", need_delim)
	dp_address, need_delim = correctInputArg(dp_address, "dp_address", need_delim)
	email, need_delim = correctInputArg(email, "cl_email", need_delim)
	return fmt.Sprintf("select * from search(%s%s%s%s%s);", order_name, cl_name, cl_surname, dp_address, email)
} 




func Search(db *sql.DB, order_name string, cl_name string, cl_surname string, email string, dp_address string)(*Utils.SearchResults) {

	str := createSelectStr(order_name, cl_name, cl_surname, email, dp_address)
	fmt.Println(str)
	rows, err := db.Query(str)

	if err != nil {
		fmt.Println("Error occurred when searching", err.Error())
		return nil
	}

	var id int
	var place int
	var name string

	var res Utils.SearchRes
	var arr Utils.SearchResults

	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&id, &place, &name)
		if err != nil {
			fmt.Println("Error occurred when scanning", err.Error())
			return nil
		}
		res.SetIdUniq(id)
		res.SetPlace(place)
		res.SetName(name)
		arr.Add(id, place, name)
	}
	return &arr
}

// func Search(db *sql.DB, order_name string, cl_name string, cl_surname string, email string, dp_address string)(*Utils.SearchRes) {

// 	str := createSelectStr(order_name, cl_name, cl_surname, email, dp_address)
// 	fmt.Println(str)
// 	rows, err := db.Query(str)

// 	if err != nil {
// 		fmt.Println("Error occurred when searching", err.Error())
// 		return nil
// 	}

// 	var id int
// 	var place int
// 	var name string

// 	var res Utils.SearchRes
// 	var arr Utils.SearchResults

// 	defer rows.Close()
// 	for rows.Next() {
// 		err = rows.Scan(&id, &place, &name)
// 		if err != nil {
// 			fmt.Println("Error occurred when scanning", err.Error())
// 			return nil
// 		}
// 		res.SetIdUniq(id)
// 		res.SetPlace(place)
// 		res.SetName(name)
// 		arr.Add(id, place, name)
// 	}
// 	return &res
// 	// return res.id_uniq, res.id_placement, res.prod_name
// }