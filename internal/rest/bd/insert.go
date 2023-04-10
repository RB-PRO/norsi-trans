package bd

import (
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Добавить запись в базу данных
func (b *Base) Add(user, pass string) error {

	// Сделать запрос к БД на добавление записи
	_, ErrorInsert := b.DB.Exec("INSERT INTO users(user, pass) VALUES(?, ?)", user, pass)
	fmt.Println(ErrorInsert)
	if ErrorInsert != nil {
		return ErrorInsert
	}

	return nil
}
