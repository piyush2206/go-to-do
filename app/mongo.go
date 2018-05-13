package app

import (
	mgo "gopkg.in/mgo.v2"
)

const (
	dbToDo = "to-do"

	collLists = "lists"
	collTasks = "tasks"
)

type (
	dbCollections struct {
		Lists *mgo.Collection
		Tasks *mgo.Collection
	}
)

func initMongoDB(ctx *Context) {
	session, err := mgo.Dial("mongo:27017")
	if err != nil {
		panic(err)
	}

	// Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	ctx.Mdb = initCollections(session)
}

func initCollections(session *mgo.Session) (collections *dbCollections) {
	todo := session.DB(dbToDo)

	collections = &dbCollections{
		Lists: todo.C(collLists),
		Tasks: todo.C(collTasks),
	}

	return
}
