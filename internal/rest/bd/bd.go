package bd

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Структура базы данных
type Base struct {
	DB *sql.DB // База данных
}

// Создать экземпляр Base
func New() (*Base, error) {

	// Открывваем базу данных
	DB, ErrorOpenDB := sql.Open("sqlite3", "gopher.db")
	if ErrorOpenDB != nil {
		return nil, ErrorOpenDB
	}

	// Создать таблицу
	_, ErrorCreateTable := DB.Exec(`CREATE TABLE IF NOT EXISTS users(
user TEXT,
pass TEXT
);`)
	if ErrorCreateTable != nil {
		return nil, ErrorCreateTable
	}

	// Пингуем БД
	ErrorPing := DB.Ping()
	if ErrorPing != nil {
		return nil, ErrorPing
	}

	return &Base{DB}, nil
}

// Закрыть БД
func (b *Base) Close() error {
	return b.DB.Close()
}
