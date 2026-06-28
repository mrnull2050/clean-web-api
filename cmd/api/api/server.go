package main

import (
	"fmt"
	"net/http"
	"time"
)

func (app *application) Serve() error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.Port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
	}
	fmt.Printf("server running on port %d", app.Port)
	return server.ListenAndServe()
	
}
