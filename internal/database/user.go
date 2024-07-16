package database

import "time"

type User struct {
	ID        int64     `db:"id" json:"id"`
	FirstName string    `db:"firstname" json:"firstname"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`

	db *DB
}

func (db *DB) HasUser(id int64) (has bool) {
	query := `SELECT EXISTS (SELECT 1 FROM users WHERE id=$1)`
	_ = db.QueryRow(query, id).Scan(&has)
	return has
}

func (db *DB) InsertUser(user User) error {
	query := `INSERT INTO users (id, firstname, created_at) VALUES ($1, $2, $3)`
	_, err := db.Exec(query, user.ID, user.FirstName, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
