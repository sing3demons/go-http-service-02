package mux

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sing3demons/go-http-service/router/ctx"
)

type myContext struct {
	w http.ResponseWriter
	r *http.Request
}

func NewMyContext(w http.ResponseWriter, r *http.Request) ctx.IContext {
	return &myContext{w, r}
}

func (ctx *myContext) JSON(code int, obj any) {
	ctx.w.Header().Set("Content-Type", "application/json; charset=UTF8")
	ctx.w.WriteHeader(code)
	json.NewEncoder(ctx.w).Encode(obj)
}

func (ctx *myContext) RequestURI() string {
	return ctx.r.RequestURI
}

func (ctx *myContext) BodyParser(obj any) error {
	decoder := json.NewDecoder(ctx.r.Body)
	decoder.UseNumber()
	decoder.DisallowUnknownFields()
	return decoder.Decode(obj)
}

func (ctx *myContext) ReadInput(obj any) error {
	body, err := io.ReadAll(ctx.r.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, &obj); err != nil {
		return err
	}
	return nil
}

func (ctx *myContext) ReadBody() ([]byte, error) {
	body, err := io.ReadAll(ctx.r.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (ctx *myContext) Param(key string) string {
	return mux.Vars(ctx.r)[key]
}

func (ctx *myContext) Query(key string) string {
	return ctx.r.URL.Query().Get(key)
}
