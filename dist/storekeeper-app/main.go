package main

import (
	"github.com/senny-matrix/celeritas"
	"github.com/senny-matrix/github.com/senny-matrix/storekeeper-app/data"
	"github.com/senny-matrix/github.com/senny-matrix/storekeeper-app/handlers"
	"github.com/senny-matrix/github.com/senny-matrix/storekeeper-app/middleware"
)

type application struct {
	App        *celeritas.Celeritas
	Handlers   *handlers.Handlers
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
