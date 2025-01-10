package middleware

import (
	"github.com/senny-matrix/celeritas"
	"github.com/senny-matrix/github.com/senny-matrix/storekeeper-app/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
