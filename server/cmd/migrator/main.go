package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"gerello/config"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	migrationsData, err := getMigrationsFilesData()
	if err != nil {
		log.Fatal(err)
	}

	err = executeMigrations(migrationsData, config.ConnStrPostgres)
	if err != nil {
		log.Fatal(err)
	}
}

func executeMigrations(migrations []byte, connStr string) error {
	pg, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}
	defer pg.Close()

	_, err = pg.Exec(string(migrations))
	if err != nil {
		return err
	}

	return nil
}

func getMigrationsFilesData() ([]byte, error) {
	var sqlQueryData []byte

	files, err := os.ReadDir("./migrations")
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		fileName := file.Name()
		if strings.Contains(fileName, "down") {
			continue
		}

		filePath := fmt.Sprintf("./migrations/%s", fileName)
		data, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		sqlQueryData = append(sqlQueryData, data...)
	}

	return sqlQueryData, nil
}
