package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/gruz0/monitoring-configuration-service/internal/configuration"
	"github.com/gruz0/monitoring-configuration-service/internal/persistence"
)

const (
	defaultConfigurationDatabaseURL = "host=localhost user=app password=password dbname=app_development sslmode=disable TimeZone=UTC"
)

func main() {
	var (
		dbURL = envString("MONITORING_CONFIGURATION_DB_URL", defaultConfigurationDatabaseURL)

		httpAddr                 = flag.String("http.addr", ":8080", "HTTP listen address")
		configurationDatabaseURL = flag.String("db.url", dbURL, "Database connection string")
	)

	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	persistenceLogger := log.With(logger, "component", "persistence")

	db, err := persistence.New(*configurationDatabaseURL)

	if err != nil {
		_ = persistenceLogger.Log(
			"method", "connect_to_database",
			"error", "Unable to connect to a database",
			"description", err.Error(),
		)

		os.Exit(1)
	}

	if err := persistence.AutoMigrate(db.DB); err != nil {
		_ = persistenceLogger.Log(
			"method", "migrate_database",
			"error", "Error while migrating a database",
			"description", err.Error(),
		)

		os.Exit(1)
	}

	if envString("MONITORING_CONFIGURATION_DB_SEED", "0") == "1" {
		if err := db.Seed(); err != nil {
			_ = persistenceLogger.Log(
				"method", "seed_database",
				"error", "Error while seeding a database",
				"description", err.Error(),
			)

			os.Exit(1)
		}
	}

	var cs configuration.Service
	cs = configuration.NewService(*db.Sites)
	cs = configuration.NewLoggingService(log.With(logger, "component", "configuration"), cs)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	mux.Handle("/configurations", configuration.MakeHandler(cs, httpLogger))

	http.Handle("/", accessControl(mux))

	errs := make(chan error, 2)

	go func() {
		_ = logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	_ = logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)

	if e == "" {
		return fallback
	}

	return e
}
