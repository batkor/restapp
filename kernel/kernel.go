package kernel

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type errHandler func(current []byte)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Database() *sql.DB {
	Settings := GetSettings()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Settings.Database.Host, Settings.Database.Port, Settings.Database.User, Settings.Database.Password, Settings.Database.Dbname)

	db, err := sql.Open("postgres", psqlInfo)
	Check(err)

	//defer func(db *sql.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		panic(err)
	//	}
	//}(db)

	err = db.Ping()
	Check(err)

	return db
}

func Bootstrap() {
	Database()
}
