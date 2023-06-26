package app

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/go-chi/chi"
)

type Container struct {
	Context context.Context
	Mux     *chi.Mux
	Server  *http.Server
}

func InitContainer(ctx context.Context) *Container {
	return &Container{
		Context: ctx,
		Mux:     chi.NewRouter(),
	}
}

func (c *Container) HttpServe(ctx context.Context, port string) {

	c.Server = &http.Server{
		Addr:    ":" + port,
		Handler: c.Mux,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}

	err := c.Server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Server started at ", port)
}
