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
	"github.com/sing3demons/go-http-service/router"
)

type microservice struct {
	*mux.Router
}

func NewMicroservice() router.IMicroservice {
	r := mux.NewRouter()
	// r.Use(loggingMiddleware)
	r.Use(handlers.RecoveryHandler(handlers.PrintRecoveryStack(true)))
	return &microservice{r}
}

// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		log.Println(r.RequestURI)
// 		next.ServeHTTP(w, r)
// 	})
// }

func (ms *microservice) Use(mwf ...router.HandlerFunc) {
	ms.Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for _, h := range mwf {
				h(NewMyContext(w, r))
			}
			next.ServeHTTP(w, r)
		})
	})
}

// func (ms *microservice) Get(path string, h router.HandlerFunc) {
// 	ms.Router.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
// 		h(NewMyContext(w, r))
// 	}).Methods(http.MethodGet)
// }

func (ms *microservice) GET(path string, handlers ...router.HandlerFunc) {
	path = regexp.MustCompile(`/:([^/]+)`).ReplaceAllString(path, "/{$1}")
	r := ms.Router.PathPrefix(path).Subrouter()
	for index := range handlers {
		if index == 0 {
			r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlers[0](NewMyContext(w, r))
			}).Methods(http.MethodGet)
		} else {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					handlers[index](NewMyContext(w, r))
					next.ServeHTTP(w, r)
				})
			})
		}
	}
	// r.Use(mwf...)
	// r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	h[0](NewMyContext(w, r))
	// }).Methods(http.MethodGet)
}

func (ms *microservice) POST(path string, handlers ...router.HandlerFunc) {
	r := ms.Router.PathPrefix(path).Subrouter()
	// r.Use(mwf...)
	// r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	h(NewMyContext(w, r))
	// }).Methods(http.MethodPost)
	for index := range handlers {
		if index == 0 {
			r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlers[0](NewMyContext(w, r))
			}).Methods(http.MethodPost)
		} else {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					handlers[index](NewMyContext(w, r))
					next.ServeHTTP(w, r)
				})
			})
		}
	}
}

func (ms *microservice) PUT(path string, handlers ...router.HandlerFunc) {
	path = regexp.MustCompile(`/:([^/]+)`).ReplaceAllString(path, "/{$1}")
	r := ms.Router.PathPrefix(path).Subrouter()
	// r.Use(mwf...)
	// r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	h(NewMyContext(w, r))
	// }).Methods(http.MethodPut)
	for index := range handlers {
		if index == 0 {
			r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlers[0](NewMyContext(w, r))
			}).Methods(http.MethodPut)
		} else {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					handlers[index](NewMyContext(w, r))
					next.ServeHTTP(w, r)
				})
			})
		}
	}
}

func (ms *microservice) PATCH(path string, handlers ...router.HandlerFunc) {
	path = regexp.MustCompile(`/:([^/]+)`).ReplaceAllString(path, "/{$1}")
	r := ms.Router.PathPrefix(path).Subrouter()
	// r.Use(mwf...)
	// r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	h(NewMyContext(w, r))
	// }).Methods(http.MethodPatch)
	for index := range handlers {
		if index == 0 {
			r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlers[0](NewMyContext(w, r))
			}).Methods(http.MethodPatch)
		} else {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					handlers[index](NewMyContext(w, r))
					next.ServeHTTP(w, r)
				})
			})
		}
	}
}

func (ms *microservice) DELETE(path string, handlers ...router.HandlerFunc) {
	path = regexp.MustCompile(`/:([^/]+)`).ReplaceAllString(path, "/{$1}")
	r := ms.Router.PathPrefix(path).Subrouter()
	// r.Use(mwf...)
	// r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	h(NewMyContext(w, r))
	// }).Methods(http.MethodDelete)
	for index := range handlers {
		if index == 0 {
			r.NewRoute().HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlers[0](NewMyContext(w, r))
			}).Methods(http.MethodDelete)
		} else {
			r.Use(func(next http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					handlers[index](NewMyContext(w, r))
					next.ServeHTTP(w, r)
				})
			})
		}
	}
}

func (ms *microservice) StartHttp() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":8080"
	}
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
	srv := &http.Server{
		Handler:      ms.Router,
		Addr:         PORT,
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
