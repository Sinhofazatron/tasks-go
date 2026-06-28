package core_http_server

import (
	"fmt"
	"net/http"

	core_hhtp_middleware "github.com/Sinhofazatron/tasks-go/internal/core/transport/http/middleware"
	core_http_middleware "github.com/Sinhofazatron/tasks-go/internal/core/transport/http/middleware"
)

type APIVersion string

var (
	APIVersion1 = APIVersion("v1")
	APIVersion2 = APIVersion("v2")
)

type APIVersionRouter struct {
	*http.ServeMux
	apiVersion APIVersion
	middleware []core_hhtp_middleware.Middleware
}

func NewAPIVersionRouter(apiVersion APIVersion, middleware ...core_hhtp_middleware.Middleware) *APIVersionRouter {
	return &APIVersionRouter{
		ServeMux:   http.NewServeMux(),
		apiVersion: apiVersion,
		middleware: middleware,
	}
}

func (r *APIVersionRouter) RegisterRouters(routers ...Route) {
	for _, route := range routers {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)

		r.Handle(pattern, route.WithMiddleware())
	}
}

func (r *APIVersionRouter) WithMiddleware() http.Handler {
	return core_http_middleware.ChainMiddleware(r, r.middleware...)
}
