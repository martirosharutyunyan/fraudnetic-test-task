package start

import (
	"context"
	"errors"
	"fmt"
	"github.com/gocql/gocql"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/config"
	"github.com/martirosharutyunyan/fraudnetic-test-task/internal/modules/controller"
	"github.com/scylladb/gocqlx/v3"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Bootstrap(env, host, port string) (<-chan error, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	ctx := context.Background()

	err := config.Load(".env")
	if err != nil {
		if env != "production" {
			log.Println("Error loading .env file")
			return nil, err
		}
	}

	cluster := gocql.NewCluster(config.GetDBHost())
	cluster.Keyspace = "fraudnetic"
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return nil, err
	}

	srv := controller.NewServer(&controller.ServerConfig{
		DB: &session,
	})

	errC := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	go func() {
		<-ctx.Done()

		log.Println("Shutdown signal received")

		ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)

		defer func() {
			stop()
			cancel()
			close(errC)
		}()

		if err := srv.ShutdownWithContext(ctxTimeout); err != nil { //nolint: contextcheck
			errC <- err
		}

		log.Println("Shutdown completed")
	}()

	go func() {
		// "ListenAndServe always returns a non-nil error. After Shutdown or Close, the returned error is
		// ErrServerClosed."
		if err := srv.Listen(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errC <- err
		}
	}()

	return errC, nil
}
