package bd

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Добавить запись в базу данных
func (b *Base) Add(user, pass string) error {

	// Сделать запрос к базе, чтобы пароверить наличие данного пользователя
	var OldUser string
	ErrorSelect := b.DB.QueryRow("SELECT user FROM users WHERE user=?", user).Scan(&OldUser)

	switch {
	case ErrorSelect == sql.ErrNoRows:
		// Добавляем новую запись в БД
		_, ErrorInsert := b.DB.Exec("INSERT INTO users(user, pass) VALUES(?, ?)", user, pass)
		if ErrorInsert != nil {
			return ErrorInsert
		}
	default:
		return ErrorSelect
	}

	return nil
}
