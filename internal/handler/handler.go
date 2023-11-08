package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"warehouseWeb/internal/html_change"
	"warehouseWeb/internal/searchStruct"
	sqlImport "warehouseWeb/internal/sql"
)

type Order struct {
	order_name string
	name       string
	surname    string
	email      string
	adress     string
}

// After login - true, before - false
var access_after_login = false

func buttonGive(w http.ResponseWriter, r *http.Request) {

}

func viewList(w http.ResponseWriter, r *http.Request) {
	if access_after_login {
		var order Order
		order.order_name = r.FormValue("order_name")
		order.name = r.FormValue("name")
		order.surname = r.FormValue("surname")
		order.email = r.FormValue("email")
		order.adress = r.FormValue("adress")

		var db *sql.DB
		db, err := sqlImport.GetDB()
		if err != nil {
			fmt.Println("Error occurred while getting access to database", err.Error())
			return
		}

		fileName := "frontend/search.html"
		listFileName := "frontend/list.html"
		// С помощью этого кода я помещаю list.html в search.html, под видом template
		// https://stackoverflow.com/questions/33984147/golang-embed-html-from-file

		var res *searchStruct.SearchResults
		if order.order_name != "" || order.name != "" || order.surname != "" || order.email != "" || order.adress != "" {
			empty := true
			res = sqlImport.Search(db, order.order_name, order.name, order.surname, order.email, order.adress)
			if *res == nil {
				empty = false
			}
			if !empty {
				html_change.WriteListNotFound(listFileName, res)
				fmt.Println("Товар не найден.")
			} else {
				html_change.WriteList(listFileName, res)
			}
		} else {
			res = searchStruct.New()
			html_change.WriteList(listFileName, res)
		}

		//html_change.WriteList(listFileName, res)
		file_list, err := template.ParseFiles(fileName, listFileName)
		// t, err := template.ParseFiles("index.html", "header.html")
		if err != nil {
			fmt.Println("Error occurred when parsing html file.", err.Error())
			return
		}

		style, err := os.ReadFile("frontend/css/success.css")
		if err != nil {
			fmt.Println("Error occured when reading CSS file.")
			return
		}

		tmplData := struct {
			Style template.CSS
		}{Style: template.CSS(style)}

		_ = file_list.ExecuteTemplate(w, fileName, nil)
		err = file_list.Execute(w, tmplData)
		if err != nil {
			fmt.Println("Error occurred when executing file.", err.Error())
			return
		}

	} else {
		login(w, r)
		fmt.Println("You are not authorized.")
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login":
		login(w, r)
	case "/submit":
		loginSubmit(w, r)
	case "/add":
		addOrder(w, r)
	case "/search":
		viewList(w, r)
	case "/buttonGive":
		buttonGive(w, r)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	access_after_login = false
	fileName := "frontend/main.html"
	file_login, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error occurred when parsing html file.", err.Error())
		return
	}
	//Reading CSS
	style, err := os.ReadFile("frontend/css/main.css")
	if err != nil {
		fmt.Println("Error occured when reading CSS file.")
		return
	}

	tmplData := struct {
		Style template.CSS
	}{Style: template.CSS(style)}

	err = file_login.Execute(w, tmplData)
	if err != nil {
		fmt.Println("Error occurred when executing file.", err.Error())
		return
	}

}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	//from URL
	loginn := r.FormValue("login")
	password := r.FormValue("password")

	//Get access to DB
	var db *sql.DB
	db, err := sqlImport.GetDB()
	if err != nil {
		fmt.Println("Error occurred while getting access to database", err.Error())
		return
	}

	//Try login
	access, err := sqlImport.GetAccessLogin(db, loginn, password)
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

		//Reading CSS
		style, err := os.ReadFile("frontend/css/success.css")
		if err != nil {
			fmt.Println("Error occured when reading CSS file.")
			return
		}

		tmplData := struct {
			Style template.CSS
		}{Style: template.CSS(style)}

		err = file.Execute(w, tmplData)
		if err != nil {
			fmt.Println("Error occurred when executing file.", err.Error())
			return
		}
		access_after_login = true
		//w.WriteHeader(http.StatusOK)
	} else {
		//w.WriteHeader(http.StatusNotFound)
		login(w, r)
		fmt.Println("There is no employee with these login and password.")
	}

}

func addOrder(w http.ResponseWriter, r *http.Request) {
	if access_after_login {
		order_name := r.FormValue("order_name")
		name := r.FormValue("name")
		surname := r.FormValue("surname")
		email := r.FormValue("email")
		adress := r.FormValue("adress")

		var db *sql.DB
		db, err := sqlImport.GetDB()
		if err != nil {
			fmt.Println("Error occurred while getting access to database", err.Error())
			return
		}

		file, err := template.ParseFiles("frontend/add.html")
		if err != nil {
			fmt.Println("Error occurred when parsing the html file for success login.", err.Error())
			return
		}

		//Reading CSS
		style, err := os.ReadFile("frontend/css/success.css")
		if err != nil {
			fmt.Println("Error occured when reading CSS file.")
			return
		}

		tmplData := struct {
			Style template.CSS
		}{Style: template.CSS(style)}

		err = file.Execute(w, tmplData)
		if err != nil {
			fmt.Println("Error occurred when executing file.", err.Error())
			return
		}

		/*
			TODO: ниже add переменную (bool) и какое-то действие сделать, например, вывести, что удалось вставить или нет
		*/
		if order_name != "" && name != "" && surname != "" && email != "" && adress != "" {
			_, err = sqlImport.AddOrderToDB(db, order_name, name, surname, email, adress)
			if err != nil {
				fmt.Println("Error occured when doing query to DB", err.Error())
				return
			}
		} else {
			return
		}
	} else {
		login(w, r)
		fmt.Println("You are not authorized.")
	}
}
