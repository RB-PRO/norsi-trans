package rest

import "log"

func Start() {

	// Настраиваем
	r, ErrorSetup := setup()
	if ErrorSetup != nil {
		log.Fatal(ErrorSetup)
	}
	defer r.DB.Close() // Закрыть БД

	// Запускаем
	r.R.Run(":8080")
}
