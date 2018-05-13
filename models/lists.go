package models

import (
	"time"

	"github.com/labstack/echo"
	"github.com/piyush2206/go-to-do/app"
	"gopkg.in/mgo.v2/bson"
)

type (
	collLists struct {
		Id   bson.ObjectId `json:"id" bson:"_id"`
		Name string        `json:"name" bson:"name"`
	}
	// ReqCreateList contains request body params
	// of POST /lists api
	ReqCreateList struct {
		Name string `json:"name"`
	}
)

// CreateList creates a to-do list in DB
func CreateList(ctx echo.Context) (err error) {
	appCtx := ctx.Get("appCtx").(*app.Context)
	reqBody := ctx.Get("req").(*ReqCreateList)

	insertLists := bson.M{
		"name":      reqBody.Name,
		"updatedAt": time.Now(),
	}

	if err = appCtx.Mdb.Lists.Insert(insertLists); err != nil {
		return
	}
	return
}

func GetLists(ctx echo.Context) (lists []collLists, err error) {
	appCtx := ctx.Get("appCtx").(*app.Context)

	listId := ctx.Param("id")

	queryLists := bson.M{}
	if listId != "" {
		queryLists = bson.M{
			"_id": bson.ObjectIdHex(listId),
		}
	}

	if err = appCtx.Mdb.Lists.Find(queryLists).All(&lists); err != nil {
		return
	}

	return
}
