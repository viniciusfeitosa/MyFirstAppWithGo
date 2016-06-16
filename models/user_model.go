package models

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// )

// const (
// 	// UserTableName is the name of the table for the user model
// 	UserTableName = "user"
// 	// UserIDCol is the column name of the model's ID
// 	UserIDCol = "id_user"
// 	// UserNameCol is the column name of the model's name
// 	UserNameCol = "name"
// 	// UserEmailCol is the column name of the model's email
// 	UserEmailCol = "email"
// )

// // User is the type to represent an user entity
// type User struct {
// 	ID    int64  `json:"id"`
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

// // NewUser return an instance of user
// func NewUser() *User {
// 	return &User{}
// }

// // Insert is the function to insert an user
// func (u *User) Insert(db *sql.DB) (sql.Result, error) {
// 	return db.Exec(
// 		fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES(?, ?)", UserTableName, UserNameCol, UserEmailCol),
// 		u.Name,
// 		u.Email,
// 	)
// }

// // SelectAll selects all users.
// func (u *User) SelectAll(db *sql.DB) ([]User, error) {
// 	rows, err := db.Query(
// 		fmt.Sprintf(
// 			"SELECT %s, %s, %s FROM %s ORDER BY %s ASC",
// 			UserIDCol,
// 			UserNameCol,
// 			UserEmailCol,
// 			UserTableName,
// 			UserNameCol,
// 		),
// 	)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users []User
// 	for rows.Next() {
// 		var retName, retEmail string
// 		var retID int64
// 		if err := rows.Scan(&retID, &retName, &retEmail); err != nil {
// 			return nil, err
// 		}
// 		log.Println(retName)
// 		users = append(users, User{ID: retID, Name: retName, Email: retEmail})
// 	}
// 	return users, nil
// }

// // SelectOne selects a user with the given ID
// func (u *User) SelectOne(db *sql.DB, id int64) error {
// 	row := db.QueryRow(
// 		fmt.Sprintf(
// 			"SELECT %s, %s, %s FROM %s WHERE %s=?",
// 			UserIDCol,
// 			UserNameCol,
// 			UserEmailCol,
// 			UserTableName,
// 			UserIDCol,
// 		),
// 		id,
// 	)
// 	var retName, retEmail string
// 	var retID int64
// 	if err := row.Scan(&retID, &retName, &retEmail); err != nil {
// 		return err
// 	}
// 	u.ID = retID
// 	u.Name = retName
// 	u.Email = retEmail
// 	return nil
// }

// // Update updates the user with the given ID
// func (u *User) Update(db *sql.DB, id int64) error {
// 	_, err := db.Exec(
// 		fmt.Sprintf(
// 			"UPDATE %s SET %s=?,%s=? WHERE %s=?",
// 			UserTableName,
// 			UserNameCol,
// 			UserEmailCol,
// 			UserIDCol,
// 		),
// 		u.Name,
// 		u.Email,
// 		id,
// 	)
// 	return err
// }

// // Delete deletes the user with the given ID
// func (u *User) Delete(db *sql.DB, id int64) error {
// 	log.Println(u.ID)
// 	_, err := db.Exec(
// 		fmt.Sprintf(
// 			"DELETE FROM %s WHERE %s=?",
// 			UserTableName,
// 			UserIDCol,
// 		),
// 		id,
// 	)
// 	return err
// }
