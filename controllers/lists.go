package controllers

import (
	"net/http"

	"github.com/piyush2206/go-to-do/models"

	"github.com/labstack/echo"
)

// CreateList is contoller for POST /lists api
func CreateList(ctx echo.Context) (err error) {
	res := formatResponse("Failed")

	defer func() {
		httpStatus := http.StatusBadRequest
		if err == nil {
			httpStatus = http.StatusOK
		}
		ctx.JSON(httpStatus, res)
	}()

	reqBody := new(models.ReqCreateList)

	if err = ctx.Bind(reqBody); err != nil {
		return
	}
	ctx.Set("req", reqBody)

	if err = models.CreateList(ctx); err != nil {
		return
	}

	res = formatResponse("Success")
	return
}