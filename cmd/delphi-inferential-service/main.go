package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/pkg/errors"

	"github.com/Max-Gabriel-Susman/delphi-inferential-service/internal/handler"
)

const (
	exitCodeErr       = 1
	exitCodeInterrupt = 2
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	defer func() {
		signal.Stop(signalChan)
		cancel()
	}()
	go func() {
		select {
		case <-signalChan: // first signal, cancel context
			cancel()
		case <-ctx.Done():
		}
		<-signalChan // second signal, hard exit
		os.Exit(exitCodeInterrupt)
	}()
	if err := run(ctx, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitCodeErr)
	}
}

func run(ctx context.Context, _ []string) error {
	var cfg struct {
		ServiceName string `env:"SERVICE_NAME" envDefault:"delphi-inferential-service"`
		Env         string `env:"ENV" envDefault:"local"`
		Database    struct {
			User   string `env:"INFERENTIAL_DB_USER,required"`
			Pass   string `env:"INFERENTIAL_DB_PASSWORD,required"`
			Host   string `env:"INFERENTIAL_DB_HOST"`
			Port   string `env:"INFERENTIAL_DB_PORT" envDefault:"3306"`
			DBName string `env:"INFERENTIAL_DB_Name" envDefault:"identity"`
			Params string `env:"INFERENTIAL_DB_Param_Overrides" envDefault:"parseTime=true"`
		}
		Datadog struct {
			Disable bool `env:"DD_DISABLE"`
		}
		Migration struct {
			Enable bool `env:"ENABLE_MIGRATE"`
		}
	}
	if err := env.Parse(&cfg); err != nil {
		return errors.Wrap(err, "parsing configuration")
	}

	h := handler.API(handler.Deps{})

	// Start API Service
	api := http.Server{
		Handler: h,
		// Addr:              "127.0.0.1:80",
		Addr:              "0.0.0.0:8080",
		ReadHeaderTimeout: 2 * time.Second,
	}

	// Make a channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	// Start listening for requests
	go func() {
		// log info about this
		serverErrors <- api.ListenAndServe()
	}()
	// Shutdown

	// logic for handling shutdown gracefully
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "server error")

	case <-ctx.Done():
		// log something

		// request a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return errors.Wrap(err, "could not stop server gracefully")
		}
	}

	return nil
}