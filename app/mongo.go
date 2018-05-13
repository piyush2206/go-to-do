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

func initMongoDB(ctx *Ctx) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	ctx.mdb = initCollections(session)
}

func initCollections(session *mgo.Session) (collections *dbCollections) {
	todo := session.DB(dbToDo)

	collections = &dbCollections{
		Lists: todo.C(collLists),
		Tasks: todo.C(collTasks),
	}

	return
}
