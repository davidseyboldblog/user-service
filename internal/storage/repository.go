package storage

import (
	"database/sql"

	"userservice/internal/adding"
	"userservice/internal/listing"
	"userservice/internal/updating"

	//Database driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	createTableStatement = "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, email TEXT, phone TEXT)"
	addUserQuery         = "INSERT INTO users (name, email, phone) VALUES(?, ?, ?)"
	getUserQuery         = "SELECT id, name, email, phone FROM users WHERE id = ?"
	updateUserQuery      = "UPDATE users SET name = ?, email = ?, phone = ? WHERE id = ?"
)

type Storage struct {
	db *sql.DB
}

func New(host string) (*Storage, error) {
	s := new(Storage)
	var err error
	s.db, err = sql.Open("sqlite3", host)
	if err != nil {
		return nil, errors.Wrap(err, "Error opening DB connection")
	}

	statement, err := s.db.Prepare(createTableStatement)
	if err != nil {
		return nil, errors.Wrap(err, "Error preparing statement")
	}
	_, err = statement.Exec()
	if err != nil {
		return nil, errors.Wrap(err, "Error creating table")
	}

	logrus.Info("Connection to database successful")
	return s, nil
}

func (s *Storage) Add(u adding.User) (int64, error) {
	tx, err := s.db.Begin()
	if err != nil {
		return 0, errors.Wrap(err, "Error opening transaction")
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(addUserQuery)
	if err != nil {
		return 0, errors.Wrap(err, "Error preparing user sql statement")
	}

	res, err := stmt.Exec(u.Name, u.Email, u.Phone)
	if err != nil {
		return 0, errors.Wrap(err, "Error executing user sql statement")
	}

	userId, err := res.LastInsertId()

	tx.Commit()
	return userId, nil
}

func (s *Storage) Update(id int64, u updating.User) error {
	tx, err := s.db.Begin()
	if err != nil {
		return errors.Wrap(err, "Error opening transaction")
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(updateUserQuery)
	if err != nil {
		return errors.Wrap(err, "Error preparing user sql statement")
	}

	_, err = stmt.Exec(u.Name, u.Email, u.Phone, id)
	if err != nil {
		return errors.Wrap(err, "Error executing user sql statement")
	}

	tx.Commit()
	return nil
}

func (s *Storage) Get(id int64) (listing.User, error) {
	var user listing.User
	err := s.db.QueryRow(getUserQuery, id).Scan(&user.ID, &user.Name, &user.Email, &user.Phone)

	if err == sql.ErrNoRows {
		//Return empty user if no rows returned
		return listing.User{}, nil
	} else if err != nil {
		return listing.User{}, errors.Wrap(err, "Error querying database")
	}

	return user, nil
}
