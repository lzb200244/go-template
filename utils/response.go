package utils

import (
	"github.com/gin-gonic/gin"
	"go-template/global/code"
	"net/http"
)

/*
Created by 斑斑砖 on 2023/8/13.
Description：
	响应体
*/

type Response struct {
	Code    code.Code   `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"msg"`
}

type Page struct {
	Results  []interface{} `json:"results,omitempty"`
	Count    int           `json:"count,omitempty"`
	Page     int           `json:"page,omitempty"`
	PageSize int           `json:"pageSize,omitempty"`
}

// ResponsePage  返回存在页数的
type ResponsePage struct {
	Code    code.Code `json:"code"`
	Message string    `json:"msg"`
	Data    *Page     `json:"data"`
}

func returnJson(ctx *gin.Context, status int, code code.Code, data interface{}, msg string) {
	ctx.JSON(
		status, Response{Code: code, Message: msg, Data: data},
	)
}
func returnPage(ctx *gin.Context, status int, code code.Code, page, pageSize, count int, msg string, results []interface{}) {
	ctx.JSON(
		status, ResponsePage{
			Code: code, Message: msg, Data: &Page{
				Page: page, PageSize: pageSize, Count: count, Results: results,
			},
		},
	)
}
func Success(ctx *gin.Context, msg string, data interface{}) {
	returnJson(ctx, http.StatusOK, code.OK, data, msg)
}
func Results(ctx *gin.Context, page, pageSize, count int, msg string, results []interface{}) {
	returnPage(
		ctx, http.StatusOK, code.OK, page, pageSize, count, msg, results,
	)
}
func Fail(ctx *gin.Context, code code.Code, msg string, data interface{}) {
	returnJson(ctx, http.StatusBadRequest, code, data, msg)
	ctx.Abort()
}
