package main

import (
	"golang-interview/database"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("no match file")
		return
	}

	database.Connect()
	 filename := os.Args[1]
	 database.RunMigration(filename)
}
