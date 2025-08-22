package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"sync"

	"mocking_http/internal/client"
	"mocking_http/internal/server"
)

func runServer(addr string, exitCh <-chan struct{}) {
	srv := http.Server{
		Addr:    addr,
		Handler: &server.Server{},
	}

	// Running server in a separate routine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				slog.Error("server closed with err", slog.String("error", err.Error()))
			}
		}
	}()

	// Handling graceful shutdown in function routine to block until all connections are closed
	// In proper code it should receive a context with deadline to avoid blocking indefinitely, but here
	// there is no processing so that isn't an issue.
	<-exitCh
	if err := srv.Shutdown(context.Background()); err != nil {
		slog.Error("server shutdown with err", slog.String("error", err.Error()))
	}
}

func runClient(serverAddr string, exitCh chan<- struct{}) {
	client := client.NewClient(serverAddr, nil)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			break
		}

		size, err := client.GetSize(line)
		if err != nil {
			slog.Error("client get size", slog.String("error", err.Error()))
			fmt.Printf("client | error | %s\n", err.Error())
		} else {
			fmt.Printf("client | size  | %d\n", size)
		}

		fmt.Print("> ")
	}
	if err := scanner.Err(); err != nil {
		slog.Error("client: reading stdin", slog.String("error", err.Error()))
	}
	close(exitCh)
}

func main() {
	var (
		wG sync.WaitGroup

		exitCh = make(chan struct{})
		port   = 8080
	)

	wG.Add(2)

	go func() {
		runServer(fmt.Sprintf(":%d", port), exitCh)
		wG.Done()
	}()

	go func() {
		runClient(fmt.Sprintf("http://localhost:%d", port), exitCh)
		wG.Done()
	}()

	// Waiting for both server and client to close gracefully
	wG.Wait()
}
