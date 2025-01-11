package middleware

import (
	"github.com/senny-matrix/celeritas"
	"github.com/senny-matrix/coolapp/data"
)

type Middleware struct {
	App    *celeritas.Celeritas
	Models data.Models
}
