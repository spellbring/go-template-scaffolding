package pkg

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
	
	"{{bootstrap_template}}/pkg/cache"
	"{{bootstrap_template}}/pkg/cache/redisrepo"
	"{{bootstrap_template}}/pkg/database"
	"{{bootstrap_template}}/pkg/database/nosql"
	"{{bootstrap_template}}/pkg/subscriber"
	"{{bootstrap_template}}/pkg/log"
	"{{bootstrap_template}}/pkg/log/logger"
	"{{bootstrap_template}}/pkg/rest"
	"{{bootstrap_template}}/pkg/router"
)

type config struct {
	appName         string
	Log             logger.Logger
	dbNoSQL         nosql.NoSQL
	dbSQL           database.IPostgresHandler
	ctxTimeout      time.Duration
	webServerPort   router.Port
	ServerHttp      router.Server
	restClient      rest.Client
	event           []subscriber.EventHandler
	redisRepository redisrepo.IRedisRepository
}

func NewConfig() *config {
	return &config{}
}

func (c *config) ContextTimeout(t time.Duration) *config {
	c.ctxTimeout = t
	return c
}

func (c *config) DbNoSQL(instance int) *config {
	db, err := database.NewDatabaseNoSQLFactory(instance)
	if err != nil {
		c.Log.Fatalln(err, "Could not make a connection to the database")
	}

	c.Log.Infof("Successfully connected to the NoSQL database")

	c.dbNoSQL = db
	return c
}

func (c *config) DbSQL(instance int) *config {
	db, err := database.NewDatabaseSQLFactory(instance)
	if err != nil {
		c.Log.Fatalln(err, "Could not make a connection to the database")
	}
	if getEnvMigrate(os.Getenv("MIGRATE")) {
		db.Migrate()
	}

	c.Log.Infof("Successfully connected to the SQL database")

	c.dbSQL = db
	return c
}

func getEnvMigrate(enrichEnv string) bool {
	if len(enrichEnv) > 0 {
		shouldMigrate, err := strconv.ParseBool(enrichEnv)
		if err != nil {
			panic(err)
		}
		return shouldMigrate
	}

	return false
}

func (c *config) InitRedisRepository(instance int) *config {
	cache, err := cache.NewRedisFactory(instance)
	if err != nil {
		c.Log.Fatalln(err, "Could not make a connection to cache")
	}

	c.Log.Infof("Successfully connected to REDIS cache")

	redisRepo := redisrepo.NewRedisRepository(cache)
	c.redisRepository = redisRepo

	return c
}

func (c *config) EventSubscriber(instance int) *config {

	notificationManagementSubscriptionConfig, notificationManagementSubscriptionName, err := subscriber.NewEventSubscriberHandlerFactory(instance, "GENERIC_SUBSCRIPTION_INSTANCE", c.Log)
	if err != nil {
		c.Log.Fatalln(err, "Could not make a connection to the notification management suscription", notificationManagementSubscriptionName)
	}
	c.event = append(c.event, notificationManagementSubscriptionConfig)
	return c
}

func (c *config) ListenSubscriberEvent() error {

	notificationManagementSubscriptionConfig := c.event[0]

	return notificationManagementSubscriptionConfig.HandleMessage(context.Background())
}

func (c *config) Name(name string) *config {
	c.appName = name
	return c
}

func (c *config) Logger(instance int) *config {
	log, err := log.NewLoggerFactory(instance)
	if err != nil {
		log.Fatalln(err)
	}

	c.Log = log
	c.Log.Infof("Successfully configured LOG")
	return c
}

func (c *config) WebServerPort(port string) *config {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		c.Log.Fatalln(err)
	}

	c.webServerPort = router.Port(p)
	return c
}

func (c *config) WebServer(instance int) *config {
	s, err := router.NewWebServerFactory(
		instance,
		c.Log,
		c.webServerPort,
		c.ctxTimeout,
	)

	if err != nil {
		c.Log.Fatalln(err)
	}

	c.Log.Infof("Successfully configured router SERVER")

	c.ServerHttp = s
	return c
}

func (c *config) RestClient(instance int) *config {
	s, err := rest.NewRestClientFactory(
		instance,
		c.Log,
		c.ctxTimeout,
	)

	if err != nil {
		c.Log.Fatalln(err)
	}

	c.Log.Infof("Successfully configured REST client")

	c.restClient = s
	return c

}

func (c *config) ListenAndHttpServe() {
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	httpServer := c.ServerHttp.GetHttpServer()

	log := c.Log

	errChannel := make(chan error, 2)

	go func() {

		if err := httpServer.ListenAndServe(); err != nil {
			log.Errorf("Error from httpServer", err)
			errChannel <- err
		}
	}()

	go func() {
		if err := c.ListenSubscriberEvent(); err != nil {
			log.Errorf("Error from listener", err)
			errChannel <- err
		}
	}()

	select {
	case <-stopChannel:
		log.Errorf("OS shutdown signal received")
	case err := <-errChannel:
		errorMessage := fmt.Sprintf("Error are ocurred: %v", err)
		log.Errorf(errorMessage)
	}

	if err := httpServer.Shutdown(ctx); err != nil {
		errorMessage := fmt.Sprintf("Error during shutdown webserver: %v", err)
		log.Errorf(errorMessage)
	}
	cancel()

}

func (c *config) HttpServe() {
	stopChannel := make(chan os.Signal, 1)
	signal.Notify(stopChannel, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	httpServer := c.ServerHttp.GetHttpServer()

	log := c.Log

	errChannel := make(chan error, 1)

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			log.Errorf("Error from httpServer", err)
			errChannel <- err
		}
	}()

	select {
	case <-stopChannel:
		log.Errorf("OS shutdown signal received")
	case err := <-errChannel:
		errorMessage := fmt.Sprintf("Error are ocurred: %v", err)
		log.Errorf(errorMessage)
	}

	if err := httpServer.Shutdown(ctx); err != nil {
		errorMessage := fmt.Sprintf("Error during shutdown webserver: %v", err)
		log.Errorf(errorMessage)
	}
	cancel()

}
