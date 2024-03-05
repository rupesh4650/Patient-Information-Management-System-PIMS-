package main

import (
	"database/sql"
	"fmt"
)

func testInsert(db *sql.DB){
  patients := []Patient{
		{
			Age:         25,
			Name:        "Ramlal Yadav",
			ImageURL:     "one.jpg",
			Address:     "123 Main St, Kathmandu, Nepal",
			DiseaseList: "Fever, Cough",
			AptDate:     "2022-02-10",
		},
		{
			Age:         35,
			Name:        "Hari Limbu",
			ImageURL:     "two.jpg",
			Address:     "New Road, Pokhara, Nepal",
			DiseaseList: "Headache, Fatigue",
			AptDate:     "2022-02-11", 
		},
	}

  for _, p := range patients {
  	err:=insertPatients(db,p)
  	if err != nil {
		  fmt.Println("Error inserting patients:", err)
		  return
	  }
	}
}
