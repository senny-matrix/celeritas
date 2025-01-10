package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"os"
	"path/filepath"
	"strings"
)

func setup(arg1, arg2 string) {
	// Load environment variables
	if arg1 != "new" && arg1 != "version" && arg1 != "help" {
		err := godotenv.Load()
		if err != nil {
			exitGracefully(err)
		}

		path, err := os.Getwd()
		if err != nil {
			exitGracefully(err)
		}

		cel.RootPath = path
		cel.DB.DatabaseType = os.Getenv("DATABASE_TYPE")

	}
}

func getDSN() string {
	dbType := cel.DB.DatabaseType
	var dsn string
	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"),
			)

		}
	} else {
		dsn = "mysql://" + cel.BuildDSN()
	}
	return dsn

}

func showHelp() {
	color.Yellow(`Available commands:
 
help- show this help
version                 - print application version
migrate                 - run all up migrations that have not been run previously
migrate down            - reverse the most recent migration 
migrate reset           - reset the database to the initial state 
make migration <name>   - create two new (up and down) migration files in the migrations folder 
make auth               - create and runs migrations for auth tables and creates models and middleware for the application
make handler <name>     - create a new stub handler file in the handlers folder
make model <name>       - create a new stub model file in the data folder
make sessions           - create a table in the database as a session store
make mail <name>		- create two starter mail templates in the mail folder
`)
}

func updateSourceFiles(path string, fi os.FileInfo, err error) error {
	// check for an error before doing anything else
	if err != nil {
		return err
	}

	// check if current file is a directory
	if fi.IsDir() {
		return nil
	}

	// only check go files
	matched, err := filepath.Match("*.go", fi.Name())
	if err != nil {
		return err
	}

	// we have a matching file
	if matched {
		// read the file
		read, err := os.ReadFile(path)
		if err != nil {
			exitGracefully(err)
		}

		newContents := strings.Replace(string(read), "myapp", appURL, -1)

		// write the changed file
		err = os.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			exitGracefully(err)
		}
	}
	return nil
}

func updateSource() {
	// walk the entire project folder, including sub-folders
	err := filepath.Walk(".", updateSourceFiles)
	if err != nil {
		exitGracefully(err)
	}
}
