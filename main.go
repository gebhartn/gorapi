package main

import "os"

var (
	listen   = os.Getenv("APP_LISTEN")
	username = os.Getenv("APP_DB_USERNAME")
	password = os.Getenv("APP_DB_PASSWORD")
	dbname   = os.Getenv("APP_DB_NAME")
)

func main() {
	app := App{}

	app.Init(username, password, dbname)
	app.Run(listen)
}
