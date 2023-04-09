package bd

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Структура базы данных
type Base struct {
	DB *sql.DB // База данных
}

// Создать экземпляр Base
func New() (*Base, error) {

	// Открывваем базу данных
	DB, ErrorOpenDB := sql.Open("mysql", "root:password@/bd.sql") //"root:password@/users"
	if ErrorOpenDB != nil {
		return nil, ErrorOpenDB
	}

	// Пингуем БД
	ErrorPing := DB.Ping()
	if ErrorPing != nil {
		return nil, ErrorPing
	}

	return &Base{DB}, nil
}

// Закрыть БД
func (DB *Base) Close() error {
	return DB.DB.Close()
}
