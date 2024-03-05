package main

import (
	"database/sql"
	"fmt"
)

func createPatientsTable(db *sql.DB) error {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS patients (
		id INT PRIMARY KEY AUTO_INCREMENT,
		age INT,
		name VARCHAR(255),
		image_url VARCHAR(255),
		address VARCHAR(255),
		disease_list VARCHAR(255),
		apt_date VARCHAR(255)
		)
		`)
	if err != nil {
		return fmt.Errorf("Error creating patients table: %v", err)
	}

	return nil
}

func insertPatients(db *sql.DB, p Patient) error{
	_, err := db.Exec(`
		INSERT INTO patients (age, name, image_url, address, disease_list, apt_date)
		VALUES (?, ?, ?, ?, ?, ?)`,
		p.Age, p.Name, p.ImageURL, p.Address, p.DiseaseList, p.AptDate)
	if err != nil {
		return fmt.Errorf("Error inserting patient: %v", err)
	}
	return nil
}

func getPatients(db *sql.DB) ([]Patient, error) {
	rows, err := db.Query(`SELECT id, age, name, image_url, address, disease_list, apt_date FROM patients`)
	if err != nil {
		return nil, fmt.Errorf("Error querying patients: %v", err)
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var p Patient
		err :=rows.Scan(&p.ID, &p.Age, &p.Name, &p.ImageURL, &p.Address, &p.DiseaseList, &p.AptDate)
		if err != nil {
			return nil, fmt.Errorf("Error scanning patient row: %v", err)
		}
		patients = append(patients, p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterating over patient rows: %v", err)
	}

	return patients, nil
}

func updatePatient(db *sql.DB, patient Patient) error {
	_, err := db.Exec(`
		UPDATE patients
		SET age=?, name=?, image_url=?, address=?, disease_list=?, apt_date=?
		WHERE id=?`,
		patient.Age, patient.Name, patient.ImageURL, patient.Address, patient.DiseaseList, patient.AptDate, patient.ID)
	if err != nil {
		return fmt.Errorf("Error updating patient: %v", err)
	}

	return nil
}
