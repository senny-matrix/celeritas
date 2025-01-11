package main

import (
	"github.com/senny-matrix/celeritas"
	"github.com/senny-matrix/coolapp/data"
	"github.com/senny-matrix/coolapp/handlers"
	"github.com/senny-matrix/coolapp/middleware"
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
