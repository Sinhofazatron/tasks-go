package core_http_server

import (
	"net/http"

	core_hhtp_middleware "github.com/Sinhofazatron/tasks-go/internal/core/transport/http/middleware"
)

type Route struct {
	Method     string
	Path       string
	Handler    http.HandlerFunc
	Middleware []core_hhtp_middleware.Middleware
}

func (r *Route) WithMiddleware() http.Handler {
	return core_hhtp_middleware.ChainMiddleware(r.Handler, r.Middleware...)
}
