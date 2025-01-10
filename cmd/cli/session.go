package main

import (
	"fmt"
	"github.com/fatih/color"
	"time"
)

func doSessionTable() error {

	dbType := cel.DB.DatabaseType

	if dbType == "mariadb" {
		dbType = "mysql"
	}

	if dbType == "postgresql" {
		dbType = "postgres"
	}

	fileName := fmt.Sprintf("%d_create_sessions_table", time.Now().UnixMicro())

	upFile := cel.RootPath + "/migrations/" + fileName + "." + dbType + ".up.sql"
	downFile := cel.RootPath + "/migrations/" + fileName + "." + dbType + ".down .sql"

	err := copyFileFromTemplate("templates/migrations/"+dbType+"_session.sql", upFile)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile([]byte("drop table if exists sessions cascade"), downFile)
	if err != nil {
		exitGracefully(err)
	}

	err = doMigrate("up", " ")
	if err != nil {
		exitGracefully(err)
	}

	color.Green("   - sessions migration created and executed")
	return nil
}
