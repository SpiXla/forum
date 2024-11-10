package database

import "fmt"

func Insert(username string, email string, password string) error {

	insertUserSQL := `INSERT INTO users (username, email, password) VALUES (?, ?, ?);`
	statement, err := Db.Prepare(insertUserSQL)
	if err != nil {
		return fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer statement.Close()

	_, err = statement.Exec(username, email, password)
	if err != nil {
		return fmt.Errorf("failed to insert data: %w", err)
	}

	return nil
}