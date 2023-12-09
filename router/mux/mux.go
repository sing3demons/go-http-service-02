package mux

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/sing3demons/go-http-service/router/ctx"
)

type microservice struct {
	*mux.Router
}

func NewMicroservice() *microservice {
	r := mux.NewRouter()
	r.Use(loggingMiddleware)
	r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))
	return &microservice{r}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

type handlerFunc func(c ctx.IContext)

func (ms *microservice) Use(mwf ...mux.MiddlewareFunc) {
	ms.Router.Use(mwf...)
}

func (ms *microservice) Group(path string) *mux.Router {
	return ms.Router.PathPrefix(path).Subrouter()
}

func (ms *microservice) Get(path string, h handlerFunc) {
	ms.Router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		h(NewMyContext(w, r))
	}).Methods(http.MethodGet)
}

func (ms *microservice) GET(path string, h handlerFunc, mwf ...mux.MiddlewareFunc) {
	path = regexp.MustCompile(`/:([^/]+)`).ReplaceAllString(path, "/{$1}")
	r := ms.Router.PathPrefix(path).Subrouter()
	r.Use(mwf...)
	r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewMyContext(w, r))
	}).Methods(http.MethodGet)
}

func (ms *microservice) POST(path string, h handlerFunc, mwf ...mux.MiddlewareFunc) {
	r := ms.Router.PathPrefix(path).Subrouter()
	r.Use(mwf...)
	r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewMyContext(w, r))
	}).Methods(http.MethodPost)
}

func (ms *microservice) PUT(path string, h handlerFunc, mwf ...mux.MiddlewareFunc) {
	path = regexp.MustCompile(`/:([^/]+)`).ReplaceAllString(path, "/{$1}")
	r := ms.Router.PathPrefix(path).Subrouter()
	r.Use(mwf...)
	r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewMyContext(w, r))
	}).Methods(http.MethodPut)
}

func (ms *microservice) PATCH(path string, h handlerFunc, mwf ...mux.MiddlewareFunc) {
	path = regexp.MustCompile(`/:([^/]+)`).ReplaceAllString(path, "/{$1}")
	r := ms.Router.PathPrefix(path).Subrouter()
	r.Use(mwf...)
	r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewMyContext(w, r))
	}).Methods(http.MethodPatch)
}

func (ms *microservice) DELETE(path string, h handlerFunc, mwf ...mux.MiddlewareFunc) {
	path = regexp.MustCompile(`/:([^/]+)`).ReplaceAllString(path, "/{$1}")
	r := ms.Router.PathPrefix(path).Subrouter()
	r.Use(mwf...)
	r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h(NewMyContext(w, r))
	}).Methods(http.MethodDelete)
}

func (ms *microservice) StartHttp() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	srv := &http.Server{
		Handler:      ms.Router,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		fmt.Printf("http listen: %s\n", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("server listen err: %v\n", err)
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown: ", err)
	}
	fmt.Println("server exited")
}
