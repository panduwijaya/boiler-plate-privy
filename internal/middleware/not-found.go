package middleware

import (
	"net/http"

	"cake-store/cake-store/internal/appctx"
	"cake-store/cake-store/internal/consts"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	rsp := appctx.NewResponse().
		WithMsgKey(consts.RespNoRouteFound).
		Generate()
	w.Header().Set("Content-Type", consts.HeaderContentTypeJSON)
	w.WriteHeader(rsp.Code)
	w.Write(rsp.Byte())
	return
}
