package app

type (
	// Ctx contains pointers to all the dependencies of the app
	Ctx struct {
		mdb *dbCollections
	}
)

// Init initialises all the dependencies of the app
// and returns a context object for dependency injection
func Init() (ctx *Ctx) {
	ctx = new(Ctx)

	initMongoDB(ctx)

	return
}
