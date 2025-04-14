package Models

import (
	"errors"
	"restapidemo/db"
	"restapidemo/utils"

	"github.com/gofrs/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

func (u *User) Save() error {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.ID = uuid.Must(uuid.NewV4())

	stmt := `INSERT INTO users (id, email, password) VALUES (?, ?, ?)`

	preparedStatement, err := db.DBConnection.Prepare(stmt)
	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	_, err = preparedStatement.Exec(u.ID, u.Email, hashedPassword)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ValidateCredentials(password string) error {

	stmt := `SELECT id, password FROM users WHERE email = ?`

	preparedStatement, err := db.DBConnection.Prepare(stmt)
	if err != nil {
		return err
	}

	defer preparedStatement.Close()

	row := preparedStatement.QueryRow(u.Email)

	var retrievedPassword string

	err = row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}

	isPasswordValid := utils.ComparePassword(retrievedPassword, password)
	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
