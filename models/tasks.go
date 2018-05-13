package models

import (
	"time"

	"github.com/labstack/echo"
	"github.com/piyush2206/go-to-do/app"
	"gopkg.in/mgo.v2/bson"
)

type (
	// CollTasks is collection struct of lists
	CollTasks struct {
		Id        bson.ObjectId `json:"id" bson:"_id"`
		ListId    bson.ObjectId `json:"listId" bson:"listId"`
		Message   string        `json:"message" bson:"message"`
		IsDelete  bool          `json:"isDelete" bson:"isDelete"`
		UpdatedAt time.Time     `json:"updatedAt" bson:"updatedAt"`
	}

	// ReqCreateTask contains request body params
	// of POST /lists api
	ReqCreateTask struct {
		ListId  string `json:"listId"`
		Message string `json:"message"`
	}
)

// CreateTasks creates task
func CreateTasks(ctx echo.Context) (err error) {
	appCtx := ctx.Get("appCtx").(*app.Context)
	reqBody := ctx.Get("req").(*ReqCreateTask)

	listId := bson.ObjectIdHex(reqBody.ListId)
	queryLists := bson.M{
		"_id":      listId,
		"isDelete": false,
	}

	list := new(CollLists)
	err = appCtx.Mdb.Lists.Find(queryLists).One(list)
	if err != nil {
		return
	}

	newTask := bson.M{
		"listId":    listId,
		"message":   reqBody.Message,
		"isDelete":  false,
		"updatedAt": time.Now(),
	}

	if err = appCtx.Mdb.Tasks.Insert(newTask); err != nil {
		return
	}

	return
}

// UpdateTask updates task message or deletes tasks
// if message if present in request body, task message is updated
// else task is deleted
func UpdateTask(ctx echo.Context) (err error) {
	appCtx := ctx.Get("appCtx").(*app.Context)
	reqBody, ok := ctx.Get("req").(*ReqCreateTask)

	taskId := bson.ObjectIdHex(ctx.Param("id"))
	updateTask := bson.M{}

	if ok {
		updateTask = bson.M{
			"$set": bson.M{
				"message":   reqBody.Message,
				"updatedAt": time.Now(),
			},
		}
	} else {
		updateTask = bson.M{
			"$set": bson.M{
				"isDelete":  true,
				"updatedAt": time.Now(),
			},
		}
	}

	if err = appCtx.Mdb.Tasks.UpdateId(taskId, updateTask); err != nil {
		return
	}

	return
}
