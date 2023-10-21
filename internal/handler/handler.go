package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	sqlImport "warehouseWeb/internal/sql"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/submit":
		loginSubmit(w, r)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	fileName := "frontend/main.html"
	file_login, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error occurred when parsing html file.", err.Error())
		return
	}
	err = file_login.ExecuteTemplate(w, "main.html", nil)
	if err != nil {
		fmt.Println("Error occurred when executing file.", err.Error())
		return
	}

}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	//from URL
	login := r.FormValue("login")
	password := r.FormValue("password")

	//Get access to DB
	var db *sql.DB
	db, err := sqlImport.GetDB()
	if err != nil {
		fmt.Println("Error occurred while getting access to database", err.Error())
		return
	}

	//Try login
	access, err := sqlImport.GetAccessLogin(db, login, password)
	if err != nil {
		fmt.Println("Error occured when doing query to DB", err.Error())
		return
	}

	//If there are login and password for employee in DB
	if access {
		file, err := template.ParseFiles("frontend/success.html")
		if err != nil {
			fmt.Println("Error occurred when parsing the html file for success login.", err.Error())
			return
		}
		err = file.ExecuteTemplate(w, "success.html", nil)
		if err != nil {
			fmt.Println("Error occurred when executing file.", err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println("There is no employee with these login and password.")
	}

}
