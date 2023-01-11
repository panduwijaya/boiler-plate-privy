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
	"cake-store/cake-store/internal/dto"
	"cake-store/cake-store/pkg/logger"
	"cake-store/cake-store/pkg/tracer"

	ucase "cake-store/cake-store/internal/ucase/contract"
)

type cakeDetail struct {
	repo repositories.Cakeer
}

func NewCakeDetail(repo repositories.Cakeer) ucase.UseCase {
	return &cakeDetail{repo: repo}
}

// Serve Cake list data
func (u *cakeDetail) Serve(dctx *appctx.Data) appctx.Response {
	var (
		param presentations.CakeQuery
		ctx   = tracer.SpanStart(dctx.Request.Context(), "ucase.detail_detail")
		lf    = logger.NewFields(
			logger.EventName("cakeDetail"),
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

	param.ID = id
	logger.Info(param.ID)
	dr, err := u.repo.FindOne(ctx, param)
	
	if err != nil {
	    tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("error find data to database: %v", err), lf...)
		return *appctx.NewResponse().WithMsgKey(consts.RespError)
	}

	if 

	logger.InfoWithContext(ctx, fmt.Sprintf("success fetch cakes to database"), lf...)
	return *appctx.NewResponse().
            WithMsgKey(consts.RespSuccess).
			WithData(dto.CakeToResponse(*dr))
}