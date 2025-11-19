package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/BadiGGH/note-app/models"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func RunDB() {
	err0 := godotenv.Load()
	if err0 != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "note_app"
	cfg.ParseTime = true
	var err error

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected!")
}

// Select note by Id
func NoteById(id int64) (models.Note, error) {
	var note models.Note
	row := db.QueryRow("SELECT * FROM note WHERE id = ?", id)
	if err := row.Scan(
		&note.ID,
		&note.Title, &note.Author, &note.Body,
		&note.CreatedAt, &note.ModifiedAt); err != nil {
		if err == sql.ErrNoRows {
			return note, fmt.Errorf("ERROR: NoteById(%d): No such a note", id)
		}
		return note, fmt.Errorf("ERROR: NoteById %d: %v", id, err)
	}
	return note, nil
}

// Select all notes
func AllNotes() ([]models.Note, error) {
	var notes []models.Note
	rows, err := db.Query("SELECT * FROM note")
	if err != nil {
		return nil, fmt.Errorf("ERROR: AllNotes(): %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var note models.Note
		if err := rows.Scan(
			&note.ID,
			&note.Title, &note.Author, &note.Body,
			&note.CreatedAt, &note.ModifiedAt); err != nil {
			return nil, fmt.Errorf("ERROR: AllNotes(): %v", err)
		}
		notes = append(notes, note)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ERROR: AllNotes(): %v", err)
	}
	return notes, nil
}

// Insert a new note to databse
func InsertNote(note models.Note) (int64, error) {
	query := `INSERT INTO note 
		(title, author, body, created_at, modified_at)
		VALUES (?,?,?,NOW(),NOW())`
	result, err := db.Exec(
		query,
		note.Title, note.Author, note.Body)
	if err != nil {
		return 0, fmt.Errorf("ERROR: InsertNote(%v): %v", note, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("ERROR: InsertNote(%v): %v", note, err)
	}
	return id, nil
}
func UpdateNote(id int64, note models.Note) (int64, error) {
	query := `UPDATE note
    SET title = ?, author = ?, body = ?, modified_at = NOW()
    WHERE id = ?;`
	result, err := db.Exec(query, note.Title, note.Author, note.Body, id)
	if err != nil {
		return 0, fmt.Errorf("ERROR: UpdateNote(%d, %v): %v", id, note, err)
	}
	numberOfUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("ERROR: UpdateNote(%d, %v): %v", id, note, err)
	}
	return numberOfUpdated, nil
}
func DeleteById(id int64) (int64, error) {
	query := `DELETE FROM note WHERE id = ?`
	result, err := db.Exec(query, id)
	if err != nil {
		return 0, fmt.Errorf("ERROR: DeleteById(%d): %v", id, err)
	}
	numberOfDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("ERROR: DeleteById(%d): %v", id, err)
	}
	return numberOfDeleted, nil
}
