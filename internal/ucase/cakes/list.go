// Package cakes
// Automatic generated
package cakes

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"cake-store/cake-store/internal/appctx"
	"cake-store/cake-store/internal/common"
	"cake-store/cake-store/internal/consts"
	"cake-store/cake-store/internal/dto"
	"cake-store/cake-store/internal/presentations"
	"cake-store/cake-store/internal/repositories"
	"cake-store/cake-store/pkg/logger"
	"cake-store/cake-store/pkg/tracer"
	"cake-store/cake-store/internal/validator"

	ucase "cake-store/cake-store/internal/ucase/contract"
)

type cakeList struct {
	repo repositories.Cakeer
}

func NewCakeList(repo repositories.Cakeer) ucase.UseCase {
	return &cakeList{repo: repo}
}

// Serve Cake list data
func (u *cakeList) Serve(dctx *appctx.Data) appctx.Response {
	var (
		param presentations.CakeQuery
		ctx   = tracer.SpanStart(dctx.Request.Context(), "ucase.list_list")
		lf    = logger.NewFields(
			logger.EventName("cakeList"),
		)
	)
    defer tracer.SpanFinish(ctx)

	err := dctx.Cast(&param)
	if err != nil {
		logger.WarnWithContext(ctx, fmt.Sprintf("error parsing query url: %v", err), lf...)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError)
	}

	err = validation.ValidateStruct(&param,
		validation.Field(&param.Page, validation.Min(int64(1))),
		validation.Field(&param.Limit, validation.Min(int64(1))),
		validation.Field(&param.StartDate, validator.ValidDateTime()),
		validation.Field(&param.EndDate, validator.ValidDateTime()),
	)

	if err != nil {
		logger.WarnWithContext(ctx, fmt.Sprintf("validation error %v", err), lf...)
		return *appctx.NewResponse().WithMsgKey(consts.RespValidationError).WithError(err)
	}

	param.Limit = common.LimitDefaultValue(param.Limit)
	param.Page = common.PageDefaultValue(param.Page)

	dr, count, err := u.repo.FindWithCount(ctx, param)
	if err != nil {
	    tracer.SpanError(ctx, err)
		logger.ErrorWithContext(ctx, fmt.Sprintf("error find data to database: %v", err), lf...)
		return *appctx.NewResponse().WithMsgKey(consts.RespError)
	}

	logger.InfoWithContext(ctx, fmt.Sprintf("success fetch cakes to database"), lf...)
	return *appctx.NewResponse().
            WithMsgKey(consts.RespSuccess).
            WithData(dto.CakesToResponse(dr)).
            WithMeta(appctx.MetaData{
                    Page:       param.Page,
                    Limit:      param.Limit,
                    TotalCount: count,
                    TotalPage:  common.PageCalculate(count, param.Limit),
            })
}