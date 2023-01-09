// Package router
package router

import (
    "context"
	"encoding/json"
	"net/http"
	"runtime/debug"

	"cake-store/cake-store/internal/appctx"
    "cake-store/cake-store/internal/bootstrap"
    "cake-store/cake-store/internal/consts"
    "cake-store/cake-store/internal/handler"
    "cake-store/cake-store/internal/middleware"
    "cake-store/cake-store/internal/ucase"
	"cake-store/cake-store/internal/ucase/cakes"
	"cake-store/cake-store/internal/repositories"
    "cake-store/cake-store/pkg/logger"
    "cake-store/cake-store/pkg/routerkit"
    "cake-store/cake-store/pkg/msgx"

    ucaseContract "cake-store/cake-store/internal/ucase/contract"
)

type router struct {
	config *appctx.Config
	router *routerkit.Router
}

// NewRouter initialize new router wil return Router Interface
func NewRouter(cfg *appctx.Config) Router {
	bootstrap.RegistryMessage()
	bootstrap.RegistryLogger(cfg)

	return &router{
		config: cfg,
		router: routerkit.NewRouter(routerkit.WithServiceName(cfg.App.AppName)),
	}
}

func (rtr *router) handle(hfn httpHandlerFunc, svc ucaseContract.UseCase, mdws ...middleware.MiddlewareFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Header.Get(consts.HeaderLanguageKey)
		if !msgx.HaveLang(consts.RespOK, lang) {
			lang = rtr.config.App.DefaultLang
			r.Header.Set(consts.HeaderLanguageKey, lang)
		}

		defer func() {
			err := recover()
			if err != nil {
				w.Header().Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
				w.WriteHeader(http.StatusInternalServerError)
				res := appctx.Response{
					Code: consts.CodeInternalServerError,
				}

				res.WithLang(lang)
				logger.Error(logger.MessageFormat("error %v", string(debug.Stack())))
				json.NewEncoder(w).Encode(res.Byte())

				return
			}
		}()

		ctx := context.WithValue(r.Context(), "access", map[string]interface{}{
			"path":      r.URL.Path,
			"remote_ip": r.RemoteAddr,
			"method":    r.Method,
		})

		req := r.WithContext(ctx)

		// validate middleware
		if !middleware.FilterFunc(w, req, rtr.config, mdws) {
			return
		}

		resp := hfn(req, svc, rtr.config)
		resp.WithLang(lang)
		rtr.response(w, resp)
	}
}

// response prints as a json and formatted string for DGP legacy
func (rtr *router) response(w http.ResponseWriter, resp appctx.Response) {
	w.Header().Set(consts.HeaderContentTypeKey, consts.HeaderContentTypeJSON)
	resp.Generate()
	w.WriteHeader(resp.Code)
	w.Write(resp.Byte())
	return
}

// Route preparing http router and will return mux router object
func (rtr *router) Route() *routerkit.Router {

    rtr.router.NotFoundHandler = http.HandlerFunc(middleware.NotFound)
	root := rtr.router.PathPrefix("/").Subrouter()
	in := root.PathPrefix("/in/").Subrouter()
	liveness := root.PathPrefix("/").Subrouter()
	inV1 := in.PathPrefix("/v1/").Subrouter()

	_ = inV1

	// open tracer setup
	bootstrap.RegistryOpenTracing(rtr.config)

    // create database session
	db := bootstrap.RegistryMultiDatabase(rtr.config.WriteDB, rtr.config.ReadDB)
    //db := bootstrap.RegistryDatabase(rtr.config.WriteDB)
	repoExample := repositories.NewCake(db)

	// use case
	healthy := ucase.NewHealthCheck()
	list := cakes.NewCakeList(repoExample)
	detail := cakes.NewCakeDetail(repoExample)
	create := cakes.NewCakeCreate(repoExample)
	delete := cakes.NewCakeDelete(repoExample)
	update := cakes.NewCakeUpdate(repoExample)

	// healthy
	liveness.HandleFunc("/liveness", rtr.handle(
		handler.HttpRequest,
		healthy,
	)).Methods(http.MethodGet)

	root.HandleFunc("/cakes", rtr.handle(
	    handler.HttpRequest,
	    list,
	)).Methods(http.MethodGet)

	root.HandleFunc("/cakes/{id:[0-9]+}", rtr.handle(
	    handler.HttpRequest,
	    detail,
	)).Methods(http.MethodGet)

	root.HandleFunc("/cakes", rtr.handle(
	    handler.HttpRequest,
	    create,
	)).Methods(http.MethodPost)

	root.HandleFunc("/cakes/{id:[0-9]+}", rtr.handle(
	    handler.HttpRequest,
	    delete,
	)).Methods(http.MethodDelete)


	root.HandleFunc("/cakes/{id:[0-9]+}", rtr.handle(
	    handler.HttpRequest,
	    update,
	)).Methods(http.MethodPut)

	// this is use case for example purpose, please delete
	//repoExample := repositories.NewExample(db)
	//el := example.NewExampleList(repoExample)
	//ec := example.NewPartnerCreate(repoExample)
	//ed := example.NewExampleDelete(repoExample)

	// TODO: create your route here

	// this route for example rest, please delete
	// example list
	//inV1.HandleFunc("/example", rtr.handle(
	//    handler.HttpRequest,
	//    el,
	//)).Methods(http.MethodGet)

	//inV1.HandleFunc("/example", rtr.handle(
	//    handler.HttpRequest,
	//    ec,
	//)).Methods(http.MethodPost)

	//inV1.HandleFunc("/example/{id:[0-9]+}", rtr.handle(
	//    handler.HttpRequest,
	//    ed,
	//)).Methods(http.MethodDelete)

	return rtr.router

}