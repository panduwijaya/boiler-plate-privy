// Package cakes
// Automatic generated
package cakes

import (
	"fmt"
	"strconv"

	"github.com/gorilla/mux"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"cake-store/cake-store/internal/appctx"
	"cake-store/cake-store/internal/consts"
	"cake-store/cake-store/internal/presentations"
	"cake-store/cake-store/internal/repositories"

	"cake-store/cake-store/pkg/logger"
	"cake-store/cake-store/pkg/tracer"

	ucase "cake-store/cake-store/internal/ucase/contract"
)

type cakeDelete struct {
	repo repositories.Cakeer
}

func NewCakeDelete(repo repositories.Cakeer) ucase.UseCase {
	return &cakeDelete{repo: repo}
}

// Serve Cake list data
func (u *cakeDelete) Serve(dctx *appctx.Data) appctx.Response {
	var (
		param presentations.CakeQuery
		ctx   = tracer.SpanStart(dctx.Request.Context(), "ucase.Delete_Delete")
		lf    = logger.NewFields(
			logger.EventName("cakeDelete"),
		)
	)
	defer tracer.SpanFinish(ctx)

	err := dctx.Cast(&param)
	if err != nil {
		logger.WarnWithContext(ctx, fmt.Sprintf("error parsing query url: %v", err), lf...)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError)
	}

	err = validation.ValidateStruct(&param,
		validation.Field(&param.ID, validation.Min(int64(1))),
	)

	if err != nil {
		logger.WarnWithContext(ctx, fmt.Sprintf("validation error %v", err), lf...)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError).WithError(err)
	}
	id, err := strconv.Atoi(mux.Vars(dctx.Request)["id"])
	logger.Info(id)
	param.ID = id

	dr, err := u.repo.Delete(ctx, param)
	logger.Info(dr)
	if err != nil {
		tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("error find data to database: %v", err), lf...)
		return *appctx.NewResponse().WithMsgKey(consts.RespError)
	}

	logger.InfoWithContext(ctx, fmt.Sprintf("success fetch cakes to database"), lf...)
	return *appctx.NewResponse().
		WithMsgKey(consts.RespSuccess)
}
