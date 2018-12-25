package mysql

import (
	"database/sql"
)

type PersonRepository struct {
	Database *sql.DB
}

type Person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (repository *PersonRepository) FindAll() []*Person {
	rows, _ := repository.Database.Query(`SELECT id, t4 name, master_id age FROM mkt_sku_ranking`)

	defer rows.Close()

	people := []*Person{}

	for rows.Next() {
		var (
			id   int
			name string
			age  int
		)

		rows.Scan(&id, &name, &age)

		people = append(people, &Person{
			Id:   id,
			Name: name,
			Age:  age,
		})
	}

	return people
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{Database: database}
}
