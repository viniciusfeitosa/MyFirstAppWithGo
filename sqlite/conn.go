package sqlite

// import (
// 	"database/sql"

// 	// I've been used the underscore only to use the driver init func
// 	_ "github.com/mxk/go-sqlite/sqlite3"
// )

// // NewConn returns a new connection with database
// func NewConn() (*sql.DB, error) {
// 	db, err := sql.Open("sqlite3", "./sqlite.db")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

// // CreateTables uses db to create a new table for models, and returns the result
// func CreateTables(db *sql.DB, query []string) error {
// 	var err error
// 	for _, q := range query {
// 		_, err = db.Exec(q)
// 	}
// 	return err
// }
