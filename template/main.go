package main

import (
	"{{bootstrap_template}}/pkg"
	"{{bootstrap_template}}/pkg/database"
	"{{bootstrap_template}}/pkg/subscriber"
	"{{bootstrap_template}}/pkg/log"
	"{{bootstrap_template}}/pkg/rest"
	"{{bootstrap_template}}/pkg/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	var app = pkg.NewConfig()

	app.ContextTimeout(10).
		Name(os.Getenv("APP_NAME")).
		DbNoSQL(database.InstanceFirestoreDB).
		Logger(log.InstanceLogrusLogger).
		DbSQL(database.InstancePostgres).
		EventSubscriber(subscriber.InstanceGooglePubSubHandler).
		RestClient(rest.InstanceRestyV2).
		WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGorillaMux).
		ListenAndHttpServe()

}
