// Package router
package router

import (
	"net/http"

	"cake-store/cake-store/internal/appctx"
	"cake-store/cake-store/internal/ucase/contract"
	"cake-store/cake-store/pkg/routerkit"
)

// httpHandlerFunc is a contract http handler for router
type httpHandlerFunc func(request *http.Request, svc contract.UseCase, conf *appctx.Config) appctx.Response

// Router is a contract router and must implement this interface
type Router interface {
	Route() *routerkit.Router
}
