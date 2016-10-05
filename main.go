package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/braintree/manners"
	"github.com/gorilla/mux"
	"github.com/tixu/Auth/bolt"
	"github.com/tixu/Auth/handlers"
	"github.com/tixu/Auth/health"
	"github.com/tixu/Auth/users"
)

const version = "1.0.0"

func main() {
	var (
		httpAddr   = flag.String("http", "0.0.0.0:80", "HTTP service address.")
		healthAddr = flag.String("health", "0.0.0.0:81", "Health service address.")
		secret     = flag.String("secret", "secret", "JWT signing secret.")
	)
	flag.Parse()

	log.Println("Starting Auth service...")
	log.Printf("Health service listening on %s", *healthAddr)
	log.Printf("HTTP service listening on %s", *httpAddr)

	errChan := make(chan error, 10)
	// starting the companion server.
	hmux := http.NewServeMux()
	hmux.HandleFunc("/healthz", health.HealthzHandler)
	hmux.HandleFunc("/readiness", health.ReadinessHandler)
	hmux.HandleFunc("/healthz/status", health.HealthzStatusHandler)
	hmux.HandleFunc("/readiness/status", health.ReadinessStatusHandler)
	healthServer := manners.NewServer()
	healthServer.Addr = *healthAddr
	healthServer.Handler = hmux

	go func() {
		errChan <- healthServer.ListenAndServe()
	}()

	mux := mux.NewRouter()
	client := bolt.NewClient()
	client.Path = "db.bolt"
	err := client.Open()
	if err != nil {
		panic(err)
	}

	userService := client.Connect().GetUserService()
	user := users.User{
		Name: "user",
		// bcrypt has for "password"
		PasswordHash: "$2a$10$KgFhp4HAaBCRAYbFp5XYUOKrbO90yrpUQte4eyafk4Tu6mnZcNWiK",
		Email:        "user@example.com",
		Role:         "wtfd",
	}
	userService.AddUser(&user)
	user = users.User{
		Name: "admin",
		// bcrypt has for "password"

		PasswordHash: "$2a$08$YFB7wzCrACOcg9IIQyhqCOyJOLNvta.IyqYqVy0i556GFhj2M2YNm",
		Email:        "user@example.com",
		Role:         "wtfd",
	}
	userService.AddUser(&user)

	mux.Handle("/login", handlers.LoginHandler(*secret, &userService))
	mux.Handle("/version", handlers.VersionHandler(version))
	adminService := client.Connect().GetAdminService()
	admin := handlers.GetAdmin(*secret, &adminService)

	mux.Handle("/admin/user", admin.Add()).Methods(http.MethodPost)
	mux.Handle("/admin/user/{id}", admin.Del()).Methods(http.MethodDelete)
	mux.Handle("/admin/user", admin.ListAll()).Methods(http.MethodGet)

	httpServer := manners.NewServer()
	httpServer.Addr = *httpAddr
	//httpServer.Handler = handlers.LoggingHandler(mux)

	go func() {
		errChan <- manners.ListenAndServe(":80", mux)
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured %v. Exiting...", s))
			health.SetReadinessStatus(http.StatusServiceUnavailable)
			httpServer.BlockingClose()
			os.Exit(0)
		}
	}
}
