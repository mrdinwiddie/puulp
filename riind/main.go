package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "pb_main"
)

type Users struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

// this function is only run when the route /hello is accessed
// 
func getUser() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	sqlStatement := `SELECT * FROM users`
	rows, err := db.Query(sqlStatement)
	fmt.Println("output", rows)

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var me string = ""
	for rows.Next() {
		var user Users
		if err = rows.Scan(&user.ID, &user.Age, &user.FirstName, &user.LastName, &user.Email); err != nil {
			log.Fatal(err)
		}
		if user.FirstName == "Dennis" {
			me = user.FirstName
		}
		fmt.Println(user)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return me
}

// TODO : test getUser and load data into the SQLite db 

func main() {
	app := pocketbase.New()
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		se.Router.GET("/hello", func(re *core.RequestEvent) error {
			var myName string = getUser()
			result := fmt.Sprintf("Hello: %s", myName)
			return re.String(200, result)
		})
		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
