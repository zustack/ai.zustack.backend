package database

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

func GetUserByID(id string) (User, error) {
	var u User
	row := DB.QueryRow(`SELECT * FROM users WHERE id = ?`, id)
	if err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
      return u, fmt.Errorf("GetUserByID: %s: no such user", id)
		}
    return u, fmt.Errorf("GetUserByID: %s: %v", id, err)
	}
	return u, nil
}

func GetUserByUsername(username string) (User, error) {
	var u User
	row := DB.QueryRow(`SELECT * FROM users WHERE username = ?`, username)
	if err := row.Scan(&u.ID, &u.Username, &u.Password, &u.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return u, fmt.Errorf("GetUserByUsername %s: no such user", username)
		}
		return u, fmt.Errorf("GetUserByUsername %s: %v", username, err)
	}
	return u, nil
}

func CreateUser(username, password string) (int64, error) {
	result, err := DB.Exec(`INSERT INTO users (username, password) VALUES (?, ?)`,
		username, password)
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("CreateUser: %v", err)
	}

	return id, nil
}
