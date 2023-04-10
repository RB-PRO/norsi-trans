package bd_test

import (
	"errors"
	"testing"

	"github.com/RB-PRO/norsi-trans/internal/rest/bd"
)

func TestNew(t *testing.T) {
	// Создаём экземпляр базы данных
	DB, ErrorDB := bd.New()
	if ErrorDB != nil {
		t.Error(ErrorDB)
	}

	// Входные данные
	AddUser := "name"
	AddPassword := "qwerty"

	// Добавить в БД данные
	ErrorAdd := DB.Add(AddUser, AddPassword)
	if ErrorAdd != nil {
		t.Error(ErrorAdd)
	}

	// Получить данные из таблицы
	GetPassword, ErrorGet := DB.Get(AddUser)
	if ErrorGet != nil {
		t.Error(ErrorGet)
	}
	if AddPassword != GetPassword { // Проверка на соответствие
		t.Error(errors.New("неправильный пароль был добавлен/получен. Получено: '" + GetPassword + "', а ожидалось: '" + AddPassword + "'"))
	}

	// Удаление данных
	ErrorDelete := DB.Delete(AddUser)
	if ErrorDelete != nil {
		t.Error(ErrorDelete)
	}

	// Ещё раз делаем запрос по старым данным и проверяем наличие ошибки
	// Если есть ошибка, то можно сказать, что удаление прошло успешно
	_, ErrorGet2 := DB.Get(AddUser)
	if ErrorGet2 == nil {
		t.Error(errors.New("После удаления найдена запись"))
	}

	// Закрываем БД
	ErrorClose := DB.Close()
	if ErrorClose != nil {
		t.Error(ErrorClose)
	}

}
