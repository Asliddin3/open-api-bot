package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// User steps
const (
	EnterFullnameStep    string = "enter_fullname"
	EnterPhoneNumberStep string = "enter_phone_number"
	RegisteredStep       string = "registered"
)

type User struct {
	TgID        int64
	TgName      string
	Fullname    *string
	PhoneNumber *string
	Step        string
	CreatedAt   *time.Time
}

type StorageI interface {
	Create(u *User) (*User, error)
	Get(id int64) (*User, error)
	GetOrCreate(tgID int64, tgName string) (*User, error)
	ChangeField(tgID int64, field, value string) error
	ChangeStep(tgID int64, step string) error
}

type storagePg struct {
	db *sql.DB
}

func NewStoragePg(db *sql.DB) StorageI {
	return &storagePg{
		db: db,
	}
}

func (s *storagePg) Create(user *User) (*User, error) {
	query := `
		INSERT INTO users(
			tg_id,
			tg_name,
			step
		) VALUES($1, $2, $3)
		RETURNING
			fullname,
			phone_number,
			created_at
	`

	err := s.db.QueryRow(
		query,
		user.TgID,
		user.TgName,
		user.Step,
	).Scan(
		&user.Fullname,
		&user.PhoneNumber,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *storagePg) Get(id int64) (*User, error) {
	var result User

	query := `
		SELECT
			tg_id,
			fullname,
			phone_number,
			step,
			created_at
		FROM users
		WHERE tg_id=$1
	`

	row := s.db.QueryRow(query, id)
	err := row.Scan(
		&result.TgID,
		&result.Fullname,
		&result.PhoneNumber,
		&result.Step,
		&result.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *storagePg) GetOrCreate(tgID int64, tgName string) (*User, error) {
	user, err := s.Get(tgID)
	if errors.Is(err, sql.ErrNoRows) {
		u, err := s.Create(&User{
			TgID:   tgID,
			TgName: tgName,
			Step:   EnterFullnameStep,
		})
		if err != nil {
			return nil, err
		}

		user = u
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *storagePg) ChangeField(tgID int64, field, value string) error {
	query := fmt.Sprintf("UPDATE users SET %s=$1 WHERE tg_id=$2", field)

	result, err := s.db.Exec(query, value, tgID)
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (s *storagePg) ChangeStep(tgID int64, step string) error {
	query := "UPDATE users SET step=$1 WHERE tg_id=$2"

	result, err := s.db.Exec(query, step, tgID)
	if err != nil {
		return err
	}

	if count, _ := result.RowsAffected(); count == 0 {
		return sql.ErrNoRows
	}

	return nil
}
