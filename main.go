package main

import (
	"lensent/db"
	"lensent/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}
