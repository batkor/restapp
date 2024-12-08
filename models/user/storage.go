package user

import (
	"batkor/restapp/kernel"
	"time"
)

func (u *User) Save() {
	queryStr := `
INSERT INTO "user" (created, login, email)
VALUES ($1, $2, $3)
RETURNING id`
	create := time.Unix(time.Now().UTC().Unix(), 0)
	err := kernel.Database().QueryRow(queryStr, create, u.Login(), u.Email()).Scan(&u.id)

	if err != nil {
		panic(err)
	}
}
