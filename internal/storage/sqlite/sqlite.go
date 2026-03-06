package sqlite

import (
	"database/sql"

	"github.com/arvindkr123/student-apis/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO STUDENTS (NAME, EMAIL, AGE) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, stmtErr := stmt.Exec(name, email, age)
	if stmtErr != nil {
		return 0, stmtErr
	}

	lastId, idErr := result.LastInsertId()
	if idErr != nil {
		return 0, idErr
	}

	return lastId, nil
}

func New(cfg *config.Config) (*Sqlite, error) {

	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	// check database connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS students (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL,
		age INTEGER
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}
