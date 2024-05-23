package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/jaredhughes1012/living_resume/svc/iam/rest"
	"github.com/rs/cors"
)

func panicIfError(err error) {
	if err != nil {
		slog.Error("Error starting service", "error", err)
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	svc, err := rest.StandardService(ctx)
	c := cors.Default()
	slog.SetLogLoggerLevel(slog.LevelDebug)

	panicIfError(err)
	panicIfError(svc.Setup(ctx, false))

	r := rest.Route(svc)
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("Starting service", "port", port)
	panicIfError(http.ListenAndServe(":"+port, c.Handler(r)))
}
