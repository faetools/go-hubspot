package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	service "github.com/faetools/{{ .Name }}/internal/app"
	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "go.uber.org/automaxprocs"
)

func main() {
	{{/*
	// AppName := env.GetAppName()
	// AppEnvironment := env.Get("CONFIG_APP_ENVIRONMENT", "")
	// AppTeam := env.GetTeam()
	// AppCapability := env.Get("CONFIG_APP_FLOW", "Undefined")
	// AppTier := env.Get("CONFIG_APP_TIER", "Undefined")
	// AppInProd, _ := strconv.ParseBool(env.Get("CONFIG_APP_IN_PROD", "0"))
	// AppVersion := env.Get("CONFIG_APP_VERSION", "development") */}}

	{{if .IsHTTPService}}
	defaultHTTPPort := os.Getenv("APP_CONFIG_HTTP_PORT")
	if defaultHTTPPort == "" {
		defaultHTTPPort = "8080"
	}
	{{ end }}

	{{/*
	// log.SetLogger(os.Stdout, AppName, AppVersion, AppTeam, AppCapability, AppTier, AppInProd)
	// log.Info("initialised logger")

	// flag.Parse()

	// if runOnlyFlag != "" {
	// 	log.Infof("Run-only flag has been set to '%s'. Only this part of the service will run.", runOnlyFlag)
	// }

	// if err := sentry.Init(sentry.ClientOptions{
	// 	Dsn:         env.Get("CONFIG_SENTRY_DSN", ""),
	// 	Release:     AppVersion,
	// 	Environment: AppEnvironment,
	// }); err != nil {
	// 	log.Fatalf("Sentry init failed: %+v", err)
	// }
	// defer sentry.Flush(time.Second * 5)
	// if !env.IsLocal() {
	// 	defer sentry.Recover()
	// }
	// log.SetSentryHook(sentry.CurrentHub().Client())

	// log.Info("initialised sentry")

	// if err := xray.Configure(xray.Config{
	// 	DaemonAddr: env.Get("CONFIG_XRAY_DAEMON", ""),
	// }); err != nil {
	// 	log.Fatalf("XRAY init failed: %+v", err)
	// }

	// if env.IsLocal() {
	// 	// This is to redirect xray to the logger when running locally
	// 	xray.SetLogger(&log.XrayLocalLogger{log.DefaultLogger})
	// }

	// privhttpsrv := app.StartInternalHTTP()

	// log.Info("started http server for liveness, readiness and metrics")

	// log.Info("initialising external http server")
	*/}}

	{{if .IsHTTPService}}
	pubhttpsrv := echo.New()
	pubhttpsrv.HideBanner = true

	{{/*
	// log.Info("applying middleware to http server")

	// HTTPTimeoutDuration := env.Get("CONFIG_APP_HTTP_HANDLER_TIMEOUT", "5s")
	// defaultHTTPTimeout, err := time.ParseDuration(HTTPTimeoutDuration)
	// if err != nil {
	// 	log.Fatalf("failed to parse CONFIG_APP_HTTP_HANDLER_TIMEOUT value", err)
	// }

	// pubhttpsrv.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
	// 	Timeout: defaultHTTPTimeout,
	// }))
	*/}}

	pubhttpsrv.Use(middleware.Logger())
	pubhttpsrv.Use(middleware.Recover())
	pubhttpsrv.Use(sentryecho.New(sentryecho.Options{Repanic: true}))

	// Default endpoint
	pubhttpsrv.GET("/ping", func(ctx echo.Context) error {
		return ctx.String(http.StatusOK, "Pong")
	})
	{{ end }}

	if err := service.RunService(
		{{ if .IsHTTPService }}pubhttpsrv,{{ end }}
	); err != nil {
		log.Fatalf("failed to init service: %+v", err)
	}

	{{/* // if env.ShouldRun("http", runOnlyFlag) { 	*/}}
	{{if .IsHTTPService}}
	go func() {
		log.Printf("starting external http server on http://localhost:%s", defaultHTTPPort)
		log.Println(pubhttpsrv.Start(fmt.Sprintf(":%s", defaultHTTPPort)))
	}()
	{{end}}

	{{/*
	// }

	// app.GracefullyShutdown(
	// 	privhttpsrv,
	// 	pubhttpsrv,
	// )
	// app.Run()
	*/}}
	Run()
}

// Run will start the application and wait for termination signals.
// Any services passed in with be closed (or shutdown gracefully see Shutdowner) when the application
// receives a termination signal.
func Run() {
	log.Println("Running app")
	{{/* // log.Info("Running app") */}}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	log.Printf("App received signal: %s", <-c)
	{{/*
	// log.Infof("App received signal: %s", <-c)

	// It takes several seconds to remove the service from the http traffic routing
	// so we wait before shutting down
	// if !env.IsLocal() {
	// 	time.Sleep(kubernetesGracePeriod)
	// }

	// log.Infof("Attempting to stop services")
	// ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	// defer cancel()

	// wg := &sync.WaitGroup{}
	// wg.Add(len(closeables))

	// for _, svc := range closeables {
	// 	func(svc Closeable) {
	// 		safe.Go(func() {
	// 			shutdown(ctx, wg, svc)
	// 		})
	// 	}(svc)
	// }

	// log.Infof("App waiting for processes to close")
	// wg.Wait()

	// log.Infof("App closed")
	*/}}

	log.Println("App closed")
}
