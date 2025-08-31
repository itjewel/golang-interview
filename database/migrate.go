package database

import (
	"log"
	"os"
	"strings"
)

func RunMigration(action string) error {
	suffix := ""
	switch action {
	case "up":
		suffix = ".up.sql"
	case "down":
		suffix = ".down.sql"
	default:
		suffix = "There is no file"
	}

	dirtory, err := os.ReadDir("migrations")
	if err != nil {
		log.Println("There is no file available")
	}

	for _, file := range dirtory {
		if !strings.HasSuffix(file.Name(), suffix) {
			continue
		}
		path := "migrations/" + file.Name()
		binarycode, err := os.ReadFile(path)
		if err != nil {
			log.Println("There is no file available")
		}
		_, err = DB.Exec(string(binarycode))
		if err != nil {
			log.Println("There is no file available")
		}
	}

	return nil
}
