package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/go-git/go-git/v5"
)

var appURL string
var appName string

func doNew(arg2 string) {
	appName = strings.ToLower(arg2)
	appURL = appName

	// sanitize the application name (convert URL to a single word)
	if strings.Contains(appName, "/") {
		exploded := strings.SplitAfter(appName, "/")
		appName = exploded[len(exploded)-1]
	}

	log.Println(" --- App Name is: ", appName, " --- ")

	// git clone the skeleton application
	color.Green("\tCloning the skeleton application ...")
	_, err := git.PlainClone("./"+appName, false, &git.CloneOptions{
		// URL:      "git@github.com:senny-matrix/celeritas-app.git",
		URL:      "https://github.com/senny-matrix/celeritas-app.git",
		Progress: os.Stdout,
		Depth:    1,
	})
	if err != nil {
		exitGracefully(err)
	}

	// remove .git directory
	//err = os.RemoveAll("./" + appName + "/.git")
	err = os.RemoveAll(fmt.Sprintf("./%s/.git", appName))
	if err != nil {
		exitGracefully(err)
	}
	// create ready to go .env file
	color.Yellow("\tCreating .env file ...")
	data, err := templateFS.ReadFile("templates/env.txt")
	if err != nil {
		exitGracefully(err)
	}

	env := string(data)
	env = strings.ReplaceAll(env, "${APP_NAME}", appName)
	env = strings.ReplaceAll(env, "${KEY}", cel.RandomString(32))

	err = copyDataToFile([]byte(env), fmt.Sprintf("./%s/.env", appName))
	if err != nil {
		exitGracefully(err)
	}
	// create a make file
	if runtime.GOOS == "windows" {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.windows", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()

		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
	} else {
		source, err := os.Open(fmt.Sprintf("./%s/Makefile.mac", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer source.Close()

		destination, err := os.Create(fmt.Sprintf("./%s/Makefile", appName))
		if err != nil {
			exitGracefully(err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			exitGracefully(err)
		}
	}

	_ = os.Remove(fmt.Sprintf("./%s/Makefile.mac", appName))
	_ = os.Remove(fmt.Sprintf("./%s/Makefile.windows", appName))

	// update the go.mod file
	color.Yellow("\tCreating go.mod file ...")
	_ = os.Remove("./" + appName + "/go.mod")
	data, err = templateFS.ReadFile("templates/go.mod.txt")
	if err != nil {
		exitGracefully(err)
	}

	mod := string(data)
	mod = strings.ReplaceAll(mod, "${APP_NAME}", appURL)

	err = copyDataToFile([]byte(mod), fmt.Sprintf("./%s/go.mod", appName))
	if err != nil {
		exitGracefully(err)
	}

	// update existing .go files with correct names/imports
	color.Yellow("\tUpdating source files ...")
	err = os.Chdir("./" + appName)
	if err != nil {
		exitGracefully(err)
	}
	updateSource()

	// run go mod tidy in the project directory
	color.Yellow("\tRunning go mod tidy ...")
	cmd := exec.Command("go", "mod", "tidy")
	err = cmd.Run()
	if err != nil {
		exitGracefully(err)
	}

	color.Green("\tApplication (%s) created successfully!", appURL)

	fmt.Println("")

	color.Green("\tTo run the application, do the following:")
	color.Green("\tcd %s", appName)
	color.Green("\tmake run")

	// open the browser
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "start", "http://localhost:4000")
	} else {
		cmd = exec.Command("open", "http://localhost:4000")
	}
}
