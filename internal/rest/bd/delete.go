package bd

import (
	_ "github.com/mattn/go-sqlite3"
)

// Удалить запись из базы данных
func (b *Base) Delete(user string) error {

	// Делаем запрос к таблице
	_, ErrSelect := b.DB.Exec("delete from users where user = '" + user + "'")
	if ErrSelect != nil {
		return ErrSelect
	}

	return nil
}
