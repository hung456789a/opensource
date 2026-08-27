package bootstrap

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Application struct {
	Env   *env
	Mongo mongo.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Mongo = NewMongoDatabase(app.Env)
	return *app
}

func (app *Application) CloseDBconnection() {
	CloseMongoDBConnection(app.Mongo)
}
