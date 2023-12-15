package controllers

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/igienator/blockchain-mysql/controllers/views"

)

type MySQLConfig struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
}

func (cfg MySQLConfig) String() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

func mySQLConnect() (*sql.DB, error) {
	cfg := MySQLConfig{
		Username: "root",
		Password: "Welcome12345",
		Database: "base1",
		Host:     "localhost",
		Port:     "3309",
	}

	db, err := sql.Open("mysql", cfg.String())
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MySQL!")

	return db, nil
}

func getUsers(db *sql.DB) ([]views.User, error) {
	query := "SELECT name, surname, email, description FROM users"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []views.User

	for rows.Next() {
		var user views.User
		err := rows.Scan(&user.Name, &user.Surname, &user.Email, &user.Description)
		if err != nil {
			return nil, err
		}
		// Append the User struct to the slice
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func ReturnUsers() ([]views.User, error) {
	db, err := mySQLConnect()
	if err != nil {
		return nil, err
	}
	users, err := getUsers(db)
	if err != nil {
		return nil, err
	}

	return users, nil
}
