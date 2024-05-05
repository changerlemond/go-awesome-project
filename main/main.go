package main

import (
	"database/sql"
	"fmt"
)

func main() {
	var db *sql.DB
	var err error
	var raw *user_raw

	db, err = connectDatabase()

	if err != nil {
		panic("데이터베이스가 연결되지 않았습니다.")
	}

	fmt.Println("DB Ready.")

	_, err = createUsersTable(db)

	if err != nil {
		panic("유저 테이블이 생성되지 않았습니다.")
	}

	fmt.Println("Table Created.")

	_, err = insertUser(db, "abc@example.com", "12345678")

	if err != nil {
		fmt.Println(err)
		panic("유저가 생성되지 않았습니다.")
	}

	fmt.Println("User Created.")

	raw, err = getUser(db)

	if err != nil {
		panic("유저를 불러오는데 실패했습니다.")
	}
	fmt.Println()
	fmt.Println("user_id", raw.user_id)
	fmt.Println("email", raw.email)
	fmt.Println("password", raw.password)
}
