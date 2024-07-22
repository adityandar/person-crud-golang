package repository

import (
	"database/sql"
	"fmt"
	"mini_project_restapi/structs"
)

func GetAllPerson(db *sql.DB) (results []structs.Person, err error) {
	sql := "SELECT * FROM person"

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var person = structs.Person{}
		err := rows.Scan(&person.ID, &person.FirstName, &person.LastName)
		if err != nil {
			return nil, err
		}

		results = append(results, person)
	}

	return results, nil
}

func InsertPerson(db *sql.DB, person structs.Person) (err error) {
	sql := "INSERT INTO person (id, first_name, last_name) VALUES ($1, $2, $3)"

	_, err = db.Exec(sql, person.ID, person.FirstName, person.LastName)

	if err != nil {
		return err
	}

	return nil
}

func UpdatePerson(db *sql.DB, person structs.Person) (err error) {
	sql := "UPDATE person SET first_name = $2, last_name = $3 WHERE id = $1"

	_, err = db.Exec(sql, person.ID, person.FirstName, person.LastName)

	if err != nil {
		return err
	}

	return nil
}

func DeletePerson(db *sql.DB, person structs.Person) (err error) {
	sql := "DELETE FROM person WHERE id = $1"

	res, err := db.Exec(sql, person.ID)

	if err != nil {
		return err
	}
	affectedRows, _ := res.RowsAffected()
	fmt.Println("affected rows:", affectedRows)

	return nil
}
