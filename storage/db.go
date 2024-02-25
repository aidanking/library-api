package storage

import (
	"database/sql"
	"log"
	"os"

	"github.com/aidanking/library-api/types"
	"github.com/go-sql-driver/mysql"
)

func Connect() *sql.DB {

	config := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "library",
	}

	db, err := sql.Open("mysql", config.FormatDSN())

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	return db
}

type AuthorRepository struct {
	DB *sql.DB
}

func (r *AuthorRepository) CreateAuthor(author *types.Author) (*types.Author, error) {
	apiUser := os.Getenv("DB_USER")
	inseryQuery := `INSERT INTO library.author(first_name, middle_name, last_name, country, created_by, updated_by) VALUES(?, ?, ?, ?, ?, ?)`

	result, insertErr := r.DB.Exec(inseryQuery, author.FirstName, author.MiddleName, author.LastName, author.Country, apiUser, apiUser)

	if insertErr != nil {
		log.Println(insertErr)
		return nil, insertErr
	}

	id, idErr := result.LastInsertId()

	if idErr != nil {
		log.Println(idErr)
		return nil, idErr
	}

	author, authorErr := r.FindAuthorById(id)

	if authorErr != nil {
		return nil, authorErr
	}

	return author, nil
}

func (r *AuthorRepository) FindAuthorById(id int64) (*types.Author, error) {

	getQuery := `SELECT id, first_name, middle_name, last_name, country FROM library.author WHERE id = ?`

	row := r.DB.QueryRow(getQuery, id)

	var author types.Author
	var middleName sql.NullString

	noRowErr := row.Scan(&author.Id, &author.FirstName, &middleName, &author.LastName, &author.Country)

	if noRowErr != nil {
		log.Println(noRowErr)
		return nil, noRowErr
	}

	handleNullMiddleName(&author, middleName)

	return &author, nil
}

func (r *AuthorRepository) FindAllAuthors() ([]types.Author, error) {
	findAllQuery := `SELECT id, first_name, middle_name, last_name, country FROM library.author`

	rows, rowsErr := r.DB.Query(findAllQuery)

	if rowsErr != nil {
		log.Println(rowsErr)
		return nil, rowsErr
	}

	defer rows.Close()

	var authors []types.Author

	for rows.Next() {
		var author types.Author
		var middleName sql.NullString

		scanErr := rows.Scan(&author.Id, &author.FirstName, &middleName, &author.LastName, &author.Country)

		if scanErr != nil {
			log.Println(scanErr)
			return nil, scanErr
		}

		handleNullMiddleName(&author, middleName)

		authors = append(authors, author)
	}

	err := rows.Err()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return authors, nil
}

func handleNullMiddleName(author *types.Author, nullString sql.NullString) {
	if nullString.Valid {
		author.MiddleName = nullString.String
	}
}

func (r *AuthorRepository) UpdateAuthor(id int64, author *types.Author) (*types.Author, error) {
	updateQuery := `UPDATE library.author
	SET first_name = ?, middle_name = ?, last_name = ?, country = ?
	WHERE id = ?
	`
	var middleName sql.NullString
	if author.MiddleName != "" {
		middleName.Valid = true
		middleName.String = author.MiddleName
	} else {
		middleName.Valid = false
	}
	_, updateErr := r.DB.Exec(updateQuery, author.FirstName, middleName, author.LastName, author.Country, id)

	if updateErr != nil {
		log.Println(updateErr)
		return nil, updateErr
	}

	updatedAuthor, updatedAuthorErr := r.FindAuthorById(id)

	if updatedAuthorErr != nil {
		log.Println(updatedAuthorErr)
		return nil, updatedAuthorErr
	}

	return updatedAuthor, nil
}

func (r *AuthorRepository) DeleteAuthor(id int64) (*types.Author, error) {

	author, authorErr := r.FindAuthorById(id)

	if authorErr != nil {
		log.Println(authorErr)
		return nil, authorErr
	}

	deleteQuery := `DELETE FROM library.author WHERE id = ?`

	_, deleteErr := r.DB.Exec(deleteQuery, id)

	if deleteErr != nil {
		log.Println(deleteErr)
		return nil, deleteErr
	}
	return author, nil
}
