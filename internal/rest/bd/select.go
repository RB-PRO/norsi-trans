package bd

import (
	_ "github.com/go-sql-driver/mysql"
)

// Получить значение из базы данных по Логину
func (b *Base) Get(user string) (string, error) {
	var Username, Password string

	// Делаем запрос к таблице
	ErrSelect := b.DB.QueryRow("SELECT user, pass FROM users WHERE user=?", user).Scan(&Username, &Password)
	if ErrSelect != nil {
		return "", ErrSelect
	}

	return Password, nil
}
