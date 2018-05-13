package app

type (
	// Context contains pointers to all the dependencies of the app
	Context struct {
		Mdb *dbCollections
	}
)

// Init initialises all the dependencies of the app
// and returns a context object for dependency injection
func Init() (ctx *Context) {
	ctx = new(Context)

	initMongoDB(ctx)

	return
}
