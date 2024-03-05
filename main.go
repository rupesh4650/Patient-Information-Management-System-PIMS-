package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Patient struct {
	ID          int
	Age         int
	Name        string
	ImageURL    string
	Address     string
	DiseaseList string
	AptDate     string
}

func main() {
	db, err := sql.Open("mysql", "rupesh:satsahib@(127.0.0.1:3306)/mydb?parseTime=true")
	if err != nil {
		fmt.Println("SQL didn't open ")
		return
	}

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			username := r.FormValue("username")
			password := r.FormValue("password")
			if username == "rupesh" && password == "kabirisgod" {
				fmt.Println("Successfully logged in")
				http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			} else {
				tmpl := template.Must(template.ParseFiles("template/error.html"))
				tmpl.Execute(w, nil)
			}
		} else {
			tmpl := template.Must(template.ParseFiles("template/index.html"))
			tmpl.Execute(w, nil)
		}
	})

	http.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
//		createPatientsTable(db)
//		testInsert(db)
		pats, err := getPatients(db)
		if err != nil {
			fmt.Println("Error while selecting * from db")
		}
		tmpl := template.Must(template.ParseFiles("template/dashboard.html"))
		tmpl.Execute(w, pats)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			patient := retrievePatient(r)
			err := insertPatients(db, patient)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error adding patient: %v", err), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		}
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		queryID := r.FormValue("id")

		id, err := strconv.Atoi(queryID)
		if err != nil {
			http.Error(w, "Invalid patient ID in the URL", http.StatusBadRequest)
			return
		}
		_, er := db.Exec(`DELETE FROM patients WHERE id = ?`, id) 
		if er != nil {
			http.Error(w, fmt.Sprintf("Error deleting patient: %v", er), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			patient := retrievePatient(r)
			fmt.Println(patient)
			err := updatePatient(db, patient)
			if err != nil {
				http.Error(w, fmt.Sprintf("Error updating patient: %v", err), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
			return
		} else {
			tmpl := template.Must(template.ParseFiles("template/update.html"))
			queryID := r.FormValue("id")
			tmpl.Execute(w, struct{ ID string }{queryID})
		}
	})

    http.ListenAndServe(":8080", nil)
    fmt.Println("running localhost with port 8080")
}

func retrievePatient(r *http.Request) Patient {
	var patient Patient
    age,err := strconv.Atoi(r.FormValue("age"))
	if err != nil {
		fmt.Println("Error retrieving age data from query ")
		return patient
	}
	id,err:=strconv.Atoi(r.FormValue("id"))
	if err != nil {
		fmt.Println("Error retrieving age data from query ")
	}
	patient.ID=id
	patient.Age=age
	patient.Name = r.FormValue("name")
	patient.ImageURL = r.FormValue("image_url")
	patient.Address = r.FormValue("address")
	patient.DiseaseList = r.FormValue("disease_list")
	patient.AptDate = r.FormValue("apt_date")

	return patient
}

