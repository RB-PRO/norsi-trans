package bd

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

// Получить значение из базы данных по Логину
func (b *Base) Get(user string) (string, error) {
	var Username, Password string

	// Делаем запрос к таблице
	ErrSelect := b.DB.QueryRow("SELECT * FROM users WHERE user=?", user).Scan(&Username, &Password)
	if ErrSelect == sql.ErrNoRows { // Если нет такой записи
		return "", errors.New("Get: not found this user")
	}
	if ErrSelect != nil { // Остальные ошибки
		return "", ErrSelect
	}

	return Password, nil
}
