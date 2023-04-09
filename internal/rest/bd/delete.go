package bd

import (
	_ "github.com/go-sql-driver/mysql"
)

// Удалить запись из базы данных
func (b *Base) Delete(user string) error {

	// Делаем запрос к таблице
	_, ErrSelect := b.DB.Query("DELETE FROM users WHERE user=?", user)
	if ErrSelect != nil {
		return ErrSelect
	}

	return nil
}
