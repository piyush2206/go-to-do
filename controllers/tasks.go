package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/piyush2206/go-to-do/models"
)

// CreateTask is contoller for POST /lists route
func CreateTask(ctx echo.Context) (err error) {
	res := formatResponse("Failed")

	defer func() {
		httpStatus := http.StatusBadRequest
		if err == nil {
			httpStatus = http.StatusOK
		}
		ctx.JSON(httpStatus, res)
	}()

	reqBody := new(models.ReqCreateTask)
	if err = ctx.Bind(reqBody); err != nil {
		return
	}

	if reqBody.ListId == "" || reqBody.Message == "" {
		err = fmt.Errorf("Invalid request")
		res = formatResponse(err.Error())
		return
	}

	ctx.Set("req", reqBody)

	if err = models.CreateTasks(ctx); err != nil {
		return
	}

	res = formatResponse("Success")
	return
}

// UpdateTask is contoller for PUT /lists/:id route
func UpdateTask(ctx echo.Context) (err error) {
	res := formatResponse("Failed")

	defer func() {
		httpStatus := http.StatusBadRequest
		if err == nil {
			httpStatus = http.StatusOK
		}
		ctx.JSON(httpStatus, res)
	}()

	if ctx.Param("id") == "" {
		return
	}

	reqBody := new(models.ReqCreateTask)
	if err = ctx.Bind(reqBody); err != nil {
		return
	}

	if reqBody.Message == "" {
		err = fmt.Errorf("Invalid request")
		res = formatResponse(err.Error())
		return
	}

	ctx.Set("req", reqBody)

	err = models.UpdateTask(ctx)
	if err != nil {
		return
	}

	res = formatResponse("Success")
	return
}

// DeleteTask is contoller for DELETE /lists/:id route
func DeleteTask(ctx echo.Context) (err error) {
	res := formatResponse("Failed")

	defer func() {
		httpStatus := http.StatusBadRequest
		if err == nil {
			httpStatus = http.StatusOK
		}
		ctx.JSON(httpStatus, res)
	}()

	if ctx.Param("id") == "" {
		return
	}

	err = models.UpdateTask(ctx)
	if err != nil {
		return
	}

	res = formatResponse("Success")
	return
}
