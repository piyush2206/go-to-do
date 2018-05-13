package models

import (
	"time"

	"github.com/labstack/echo"
	"github.com/piyush2206/go-to-do/app"
	"gopkg.in/mgo.v2/bson"
)

type (
	// CollLists is collection struct of lists
	CollLists struct {
		Id        bson.ObjectId `json:"id" bson:"_id"`
		Name      string        `json:"name" bson:"name"`
		IsDelete  bool          `json:"isDelete" bson:"isDelete"`
		UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
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
		"isDelete":  false,
		"updatedAt": time.Now(),
	}

	if err = appCtx.Mdb.Lists.Insert(insertLists); err != nil {
		return
	}

	return
}

// GetLists fetches all lists or a single list if id is passed
func GetLists(ctx echo.Context) (lists []CollLists, err error) {
	appCtx := ctx.Get("appCtx").(*app.Context)

	listId := ctx.Param("id")

	queryLists := bson.M{"isDelete": false}
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

func DeleteList(ctx echo.Context) (err error) {
	appCtx := ctx.Get("appCtx").(*app.Context)

	listId := bson.ObjectIdHex(ctx.Param("id"))
	updateLists := bson.M{
		"$set": bson.M{
			"isDelete": true,
		},
	}

	if err = appCtx.Mdb.Lists.UpdateId(listId, updateLists); err != nil {
		return
	}

	return
}
